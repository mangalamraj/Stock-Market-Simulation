"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";
const Signup = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const router = useRouter();
  const handleSignup = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/signup", {
      method: "POST",
      body: JSON.stringify({ name, email, password, confirmPassword }),
    });
    if (!response.ok) {
      const error = await response.json();
      console.log(error);
      return;
    }
    const data = await response.json();
    console.log(data);
    router.push("/login");
  };

  return (
    <div>
      <h1 className="container m-auto text-center text-2xl font-bold p-4">
        Signup
      </h1>
      <form
        className="container m-auto flex flex-col gap-4 w-1/2 p-4"
        onSubmit={handleSignup}
      >
        <Input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <Input
          type="text"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <Input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <Input
          type="password"
          placeholder="Confirm Password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />
        <Button type="submit">Signup</Button>
      </form>
      <div className="container m-auto text-center text-sm font-bold text-gray-600">
        Already have an account?{" "}
        <Link href="/login" className="text-blue-500">
          Login
        </Link>
      </div>
    </div>
  );
};

export default Signup;
