import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Link from "next/link";

const Signup = () => {
  return (
    <div>
      <h1 className="container m-auto text-center text-2xl font-bold p-4">
        Signup
      </h1>
      <form className="container m-auto flex flex-col gap-4 w-1/2 p-4">
        <Input type="text" placeholder="Name" />
        <Input type="text" placeholder="Email" />
        <Input type="password" placeholder="Password" />
        <Input type="password" placeholder="Confirm Password" />
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
