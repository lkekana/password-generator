import React from "react";
import { useMutation } from "@tanstack/react-query";
import { Button } from "./ui/button";
import { CircleAlert, Clipboard, Check } from "lucide-react";
import { Spinner } from "./ui/spinner";

interface ClipboardButtonProps {
	textToCopy: string;
	className?: string;
}

const ClipboardButton: React.FC<ClipboardButtonProps> = ({
	textToCopy,
	className,
}) => {
	const handleCopy = async () => {
		await navigator.clipboard.writeText(textToCopy);
	};

	const { mutate, isPending, isError, isSuccess, reset } = useMutation({
		mutationFn: handleCopy,
		onSuccess: () => {
			setTimeout(() => reset(), 2000);
		},
	});

	return (
		<Button size="icon" className={className} onClick={() => mutate()}>
			{/* <Clipboard /> */}
			{isPending ? (
				<Spinner />
			) : isError ? (
				<CircleAlert />
			) : isSuccess ? (
				<Check />
			) : (
				<Clipboard />
			)}
		</Button>
	);
};

export default ClipboardButton;
