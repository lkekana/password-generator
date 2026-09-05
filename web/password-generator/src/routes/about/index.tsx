import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export const Route = createFileRoute("/about/")({
	component: RouteComponent,
});

function RouteComponent() {
	return (
		<div className="pt-8 w-screen flex justify-center">
		<Card className="w-full max-w-sm">
			<CardHeader>
				<CardTitle>passwords-generator</CardTitle>
				<CardDescription>
				</CardDescription>
			</CardHeader>
			<CardContent>
			</CardContent>
			<CardFooter className="flex-col gap-2">
            </CardFooter>
		</Card>
		</div>
	);
}
