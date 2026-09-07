import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Checkbox } from "#/components/ui/checkbox";
import { Marker } from "#/components/ui/marker";
import { Clipboard, Minus, Plus, RefreshCw } from "lucide-react";
import { Textarea } from "#/components/ui/textarea";
import { Slider } from "#/components/ui/slider";
import { useCallback, useEffect, useState } from "react";
import type { CheckedState } from "@radix-ui/react-checkbox";
import { generatePassword } from "#/generator";
import { ScrollArea } from "@radix-ui/react-scroll-area";
import ClipboardButton from "#/components/ClipboardButton";
import { toast } from "sonner";

export const Route = createFileRoute("/")({ component: Home });

type MyTabs = "single" | "multiple";
// const samplePasswords = [
// 	"GeneratedPassword123!",
// 	"AnotherPassword456!",
// 	"YetAnotherPassword789!",
// 	"PasswordExample000!",
// 	"FinalPasswordExample999!",
// ];

function Home() {
	const [activeTab, setActiveTab] = useState<MyTabs>("single");
	const [password, setPassword] = useState<string>("");
	const [passwordLength, setPasswordLength] = useState(16);
	const [includeUppercase, setIncludeUppercase] = useState<CheckedState>(true);
	const [includeLowercase, setIncludeLowercase] = useState<CheckedState>(true);
	const [includeNumbers, setIncludeNumbers] = useState<CheckedState>(true);
	const [includeSpecial, setIncludeSpecial] = useState<CheckedState>(false);
	const [generationDisabled, setGenerationDisabled] = useState(false);

	const [multiplePasswordCount, setMultiplePasswordCount] = useState(4);
	const [generatedPasswords, setGeneratedPasswords] = useState<string[]>([]);

	const generatePasswordSafe = useCallback(
		(
			length: number,
			includeUppercase: boolean,
			includeLowercase: boolean,
			includeNumbers: boolean,
			includeSpecialChars: boolean,
		): string => {
			// if (
			// 	!includeUppercase &&
			// 	!includeLowercase &&
			// 	!includeNumbers &&
			// 	!includeSpecialChars
			// ) {
			// 	return "";
			// }

			try {
				return generatePassword(
					length,
					includeUppercase,
					includeLowercase,
					includeNumbers,
					includeSpecialChars,
				);
			} catch (error) {
				console.error("Error generating password:", error);
				toast.error(`Error generating password: ${error}`);
				return "";
			}
		},
		[],
	);

	const handleRegenerate = () => {
		if (activeTab === "single") {
			setPassword(() =>
				generatePasswordSafe(
					passwordLength,
					!!includeUppercase,
					!!includeLowercase,
					!!includeNumbers,
					!!includeSpecial,
				),
			);
		} else if (activeTab === "multiple") {
			setGeneratedPasswords(() => {
				const newPasswords: string[] = [];
				for (let i = 0; i < multiplePasswordCount; i++) {
					newPasswords.push(
						generatePasswordSafe(
							passwordLength,
							!!includeUppercase,
							!!includeLowercase,
							!!includeNumbers,
							!!includeSpecial,
						),
					);
				}
				return newPasswords;
			});
		}
	};

	useEffect(() => {
		// setGeneratedPasswords(samplePasswords.slice(0, multiplePasswordCount));
		if (
			!includeUppercase &&
			!includeLowercase &&
			!includeNumbers &&
			!includeSpecial
		) {
			setGenerationDisabled(true);
			setPassword("");
			setGeneratedPasswords([]);
			return;
		} else {
			setGenerationDisabled(false);
		}

		if (activeTab === "single") {
			setPassword(() =>
				generatePasswordSafe(
					passwordLength,
					!!includeUppercase,
					!!includeLowercase,
					!!includeNumbers,
					!!includeSpecial,
				),
			);
		} else if (activeTab === "multiple") {
			setGeneratedPasswords(() => {
				const newPasswords: string[] = [];
				for (let i = 0; i < multiplePasswordCount; i++) {
					newPasswords.push(
						generatePasswordSafe(
							passwordLength,
							!!includeUppercase,
							!!includeLowercase,
							!!includeNumbers,
							!!includeSpecial,
						),
					);
				}
				return newPasswords;
			});
		}
	}, [
		multiplePasswordCount,
		passwordLength,
		includeUppercase,
		includeLowercase,
		includeNumbers,
		includeSpecial,
		activeTab,
		generatePasswordSafe,
	]);

	return (
		<div className="pt-6 md:pt-8 px-4 w-screen flex justify-center">
			<Tabs
				value={activeTab}
				onValueChange={(value) => setActiveTab(value as MyTabs)}
				className="max-w-160"
			>
				<TabsList className="grid w-full grid-cols-2">
					<TabsTrigger value="single">Single</TabsTrigger>
					<TabsTrigger value="multiple">Multiple</TabsTrigger>
				</TabsList>
				<Card className="gap-2 md:gap-4">
					{/* <CardHeader>
							<CardTitle>Config</CardTitle>
							<CardDescription>
								Make changes to your account here. Click save when you&apos;re
								done.
							</CardDescription>
						</CardHeader> */}
					<CardContent className="grid w-full">
						<div className="flex flex-col md:flex-row gap-2 md:gap-3 w-full md:items-center">
							<Label htmlFor="password-length">Password Length</Label>
							<div className="flex items-center gap-2 w-full">
								<Slider
									step={1}
									min={4}
									max={128}
									value={[passwordLength]}
									onValueChange={(value) => setPasswordLength(value[0])}
								/>
								<Input
									id="password-length"
									type="number"
									min={4}
									max={128}
									value={passwordLength}
									onChange={(e) => setPasswordLength(Number(e.target.value))}
									className="w-1/5 md:w-1/7 text-center px-2"
								/>
							</div>
						</div>
						<div className="flex flex-col md:flex-row pt-4 gap-4">
							<div className="flex gap-2 items-center">
								<Checkbox
									id="use-uppercase"
									checked={includeUppercase}
									onCheckedChange={(checked) => setIncludeUppercase(checked)}
								/>
								<Label htmlFor="use-uppercase" className="font-mono">
									Uppercase (ABC)
								</Label>
							</div>
							<div className="flex gap-2 items-center">
								<Checkbox
									id="use-lowercase"
									checked={includeLowercase}
									onCheckedChange={(checked) => setIncludeLowercase(checked)}
								/>
								<Label htmlFor="use-lowercase" className="font-mono">
									Lowercase (abc)
								</Label>
							</div>
							<div className="flex gap-2 items-center">
								<Checkbox
									id="use-numbers"
									checked={includeNumbers}
									onCheckedChange={(checked) => setIncludeNumbers(checked)}
								/>
								<Label htmlFor="use-numbers" className="font-mono">
									Numbers (123)
								</Label>
							</div>
							<div className="flex gap-2 items-center">
								<Checkbox
									id="use-special"
									checked={includeSpecial}
									onCheckedChange={(checked) => setIncludeSpecial(checked)}
								/>
								<Label htmlFor="use-special" className="font-mono">
									Special Characters
								</Label>
							</div>
						</div>
					</CardContent>
					<CardFooter>
						<div className="w-full flex flex-col items-center gap-4">
							<Marker variant="border" />
							<TabsContent value="single" className="w-full">
								<div className="w-full flex">
									<Input
										id="generated-password"
										type="text"
										className="grow flex-1 mr-2 font-mono font-medium"
										readOnly
										value={password}
										onChange={(e) => setPassword(e.target.value)}
									/>
									<ClipboardButton textToCopy={password} className="shrink" />
								</div>
							</TabsContent>
							<TabsContent value="multiple" className="w-full">
								<div className="flex flex-col w-full gap-4">
									<div className="flex flex-col gap-4">
										<div className="flex justify-between items-center gap-4">
											<Label htmlFor="password-count" className="w-max">
												How many passwords?
											</Label>
											<div className="flex items-center gap-2">
												<Button
													size="icon"
													variant="neutral"
													onClick={
														() =>
															setMultiplePasswordCount((prev) =>
																Math.max(1, prev - 1),
															)
														// setGeneratedPasswords((prev) =>
														// 	samplePasswords.slice(0, Math.max(1, prev.length - 1)),
														// )
													}
													disabled={generatedPasswords.length <= 1}
												>
													<Minus className="h-4 w-4" />
												</Button>

												<Input
													id="password-count"
													type="number"
													value={multiplePasswordCount}
													min={1}
													onChange={(e) => {
														const val = parseInt(e.target.value, 10);
														if (!Number.isNaN(val)) {
															// setValue(Math.min(max, Math.max(min, val)));
															setMultiplePasswordCount(() => Math.max(1, val));
															// setGeneratedPasswords((prev) => {
															// 	return samplePasswords.slice(0, Math.max(1, val));
															// });
														}
													}}
													className="w-16 text-center [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
												/>

												<Button
													size="icon"
													variant="neutral"
													onClick={
														() => setMultiplePasswordCount((prev) => prev + 1)
														// setGeneratedPasswords((prev) => {
														// 	const newCount = prev.length + 1;
														// 	return samplePasswords.slice(0, newCount);
														// })
													}
													// disabled={value >= max}
												>
													<Plus className="h-4 w-4" />
												</Button>
											</div>
										</div>

										<Textarea
											// readOnly
											rows={5}
											className="font-mono font-medium"
											value={generatedPasswords.join("\n")}
											onChange={(e) =>
												setGeneratedPasswords(
													e.target.value.split("\n").filter(Boolean),
												)
											}
										/>
									</div>
								</div>
							</TabsContent>
							<div className="w-full flex gap-2">
								<Button
									size="icon"
									className="grow"
									onClick={handleRegenerate}
									disabled={generationDisabled}
								>
									Regenerate
									<RefreshCw />
								</Button>
								{activeTab === "multiple" && (
									<ClipboardButton
										textToCopy={generatedPasswords.join("\n")}
										className="shrink"
									/>
								)}
							</div>
						</div>
					</CardFooter>
				</Card>
			</Tabs>
		</div>
	);
}
