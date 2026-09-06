import { Link, useMatchRoute } from "@tanstack/react-router";
import { Button } from "./ui/button";
import Star13 from "./stars/s13";

export function Navbar() {
	const matchRoute = useMatchRoute();
	const isHomeActive = !!matchRoute({ to: "/" });
	const isAboutActive = !!matchRoute({ to: "/about" });

	return (
		<header className="w-full border-b-4 border-black bg-white px-4 md:px-8 py-4 flex items-center justify-between shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] sticky top-0 z-50">
			<div className="flex items-center gap-2 md:gap-4">
			<Star13 className="w-16 h-16" />
			<div className="flex flex-col">
				<Link
					to="/"
					className="text-xl md:text-2xl font-extrabold tracking-tight hover:text-gray-700 transition-colors"
				>
					Passwords Generator
				</Link>
				<p className="text-xs md:text-sm text-gray-600 hidden sm:block">
					A simple, secure password generator
				</p>
			</div>
			</div>

			<nav className="flex items-center gap-2 md:gap-4">
				<Button asChild variant={isHomeActive ? "default" : "neutral"}>
					<Link
						to="/"
						activeOptions={{ exact: true }} // Crucial: prevents "/" from being active when on "/about"
						// activeProps={{
						// 	className:
						// 		"bg-black text-white shadow-none translate-x-[2px] translate-y-[2px]",
						// }}
						// inactiveProps={{
						// 	className:
						// 		"bg-white text-black shadow-[2px_2px_0px_0px_rgba(0,0,0,1)] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none",
						// }}
					>
						Home
					</Link>
				</Button>

				<Button asChild variant={isAboutActive ? "default" : "neutral"}>
					<Link
						to="/about"
						activeOptions={{ exact: true }}
						// activeProps={{
						// 	className:
						// 		"bg-black text-white shadow-none translate-x-[2px] translate-y-[2px]",
						// }}
						// inactiveProps={{
						// 	className:
						// 		"bg-white text-black shadow-[2px_2px_0px_0px_rgba(0,0,0,1)] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none",
						// }}
					>
						About
					</Link>
				</Button>

				<Button asChild variant="neutral">
					<a
						href="https://github.com/lkekana"
						target="_blank"
						rel="noopener noreferrer"
						className="p-2 border-2 border-black flex items-center justify-center"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							strokeWidth="2"
							strokeLinecap="round"
							strokeLinejoin="round"
							className="size-5"
						>
							<title>GitHub</title>
							<path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4" />
							<path d="M9 18c-4.51 2-5-2-7-2" />
						</svg>
					</a>
				</Button>
			</nav>
		</header>
	);
}
