import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function Home() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center align-middle min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)] h-screen">
      <div className="flex flex-col items-center justify-center ">
        <h1 className="text-4xl font-bold">Stock Market Simulation</h1>
        <p className="text-lg text-gray-600">
          Welcome to the stock market simulation.
        </p>
        <Button>
          <Link href="/login">Login</Link>
        </Button>
      </div>
    </div>
  );
}
