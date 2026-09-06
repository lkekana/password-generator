const UPPERCASE_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
const LOWERCASE_CHARS = "abcdefghijklmnopqrstuvwxyz";
const NUMBER_CHARS = "0123456789";
const SPECIAL_CHARS = "!@#$%^&*()-_=+[]{}|;:,.<>?/";

export function secureRandomInt(min: number, max: number): number {
	const range = max - min + 1;
	const maxSafeValue = 0x100000000 - (0x100000000 % range);
	const array = new Uint32Array(1);

	do {
		globalThis.crypto.getRandomValues(array);
	} while (array[0] >= maxSafeValue);

	return min + (array[0] % range);
}

export function secureRandomIntArray(
	min: number,
	max: number,
	length: number,
): number[] {
	const range = max - min + 1;
	const maxSafeValue = 0x100 - (0x100 % range);
	const array = new Uint8Array(length);
	globalThis.crypto.getRandomValues(array);

	// rejection sampling to avoid modulo bias
	const result: number[] = [];
	for (let i = 0; i < length; i++) {
		const fixerArray = new Uint8Array(1);
		while (array[i] >= maxSafeValue) {
			globalThis.crypto.getRandomValues(fixerArray);
			array[i] = fixerArray[0];
		}
		result.push(min + (array[i] % range));
	}
	return result;
}

function meetsRequirements(
	pwd: string,
	upper: boolean,
	lower: boolean,
	num: boolean,
	special: boolean,
): boolean {
	let hasUpper = false;
	let hasLower = false;
	let hasNum = false;
	let hasSpecial = false;

	for (const c of pwd) {
		if (c >= "A" && c <= "Z") {
			hasUpper = true;
		} else if (c >= "a" && c <= "z") {
			hasLower = true;
		} else if (c >= "0" && c <= "9") {
			hasNum = true;
		} else {
			hasSpecial = true;
		}
	}

	return (
		(!upper || hasUpper) &&
		(!lower || hasLower) &&
		(!num || hasNum) &&
		(!special || hasSpecial)
	);
}

export function generatePassword(
	length: number,
	includeUppercase: boolean,
	includeLowercase: boolean,
	includeNumbers: boolean,
	includeSpecialChars: boolean,
): string {
	let charset = "";
	if (includeUppercase) {
		charset += UPPERCASE_CHARS;
	}
	if (includeLowercase) {
		charset += LOWERCASE_CHARS;
	}
	if (includeNumbers) {
		charset += NUMBER_CHARS;
	}
	if (includeSpecialChars) {
		charset += SPECIAL_CHARS;
	}

	if (charset.length === 0) {
		throw new Error("No character sets selected");
	}

	let requiredCount = 0;
	if (includeUppercase) {
		requiredCount++;
	}
	if (includeLowercase) {
		requiredCount++;
	}
	if (includeNumbers) {
		requiredCount++;
	}
	if (includeSpecialChars) {
		requiredCount++;
	}

	if (length < requiredCount) {
		throw new Error(
			`Password length is too short for the selected character requirements. Minimum length required: ${requiredCount}`,
		);
	}

	do {
		const password = secureRandomIntArray(0, charset.length - 1, length)
			.map((i) => charset[i])
			.join("");
		if (
			meetsRequirements(
				password,
				includeUppercase,
				includeLowercase,
				includeNumbers,
				includeSpecialChars,
			)
		) {
			return password;
		}
		// biome-ignore lint/correctness/noConstantCondition: I want to loop until a valid password is generated
	} while (true);
}
