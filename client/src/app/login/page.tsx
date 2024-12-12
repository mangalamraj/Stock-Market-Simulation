"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { useState } from "react";
import { useRouter } from "next/navigation";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/user/login", {
      method: "POST",
      body: JSON.stringify({ email, password }),
    });
    if (!response.ok) {
      const error = await response.json();
      console.log(error);
      return;
    }
    const data = await response.json();
    localStorage.setItem("token", email);
    router.push("/");
    console.log(data);
  };

  return (
    <div>
      <h1 className="container m-auto text-center text-2xl font-bold p-4">
        Login
      </h1>
      <form
        className="container m-auto flex flex-col gap-4 w-1/2 p-4"
        onSubmit={handleLogin}
      >
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
        <Button type="submit">Login</Button>
      </form>
      <div className="container m-auto text-center text-sm font-bold text-gray-600">
        Donot have an account?{" "}
        <Link href="/signup" className="text-blue-500">
          Register
        </Link>
      </div>
    </div>
  );
};

export default Login;
