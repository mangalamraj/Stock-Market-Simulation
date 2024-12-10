import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Link from "next/link";

const Login = () => {
  return (
    <div>
      <h1 className="container m-auto text-center text-2xl font-bold p-4">
        Login
      </h1>
      <form className="container m-auto flex flex-col gap-4 w-1/2 p-4">
        <Input type="text" placeholder="Username" />
        <Input type="password" placeholder="Password" />
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
