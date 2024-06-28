"use client";
import { signIn, signOut, useSession } from "next-auth/react";
import Link from "next/link";
import { usePathname } from "next/navigation";

import React from "react";

const AuthButton = () => {
  const { data: session, status } = useSession();
  //   const pathname = usePathname();

  if (status === "loading") {
    return <button>Loading...</button>;
  }

  if (session) {
    return (
      <>
        <button onClick={() => signOut()}>Sign out</button>
      </>
    );
  }

  return (
    <>
      <button onClick={() => signIn()}>Sign in</button>
    </>
  );
};

const Navbar = () => {
  return (
    <header className="w-full p-4 bg-blue-600 text-white text-center">
      <div className="container mx-auto">
        <div className="flex justify-between items-center">
          <Link href={"/"}>
            <h1 className="text-1xl md:text-2xl lg:text-3xl">GIXEL</h1>
          </Link>
          <ul className="flex gap-6">
            <Link href={"/downloads"}>
              <li>Downloads</li>
            </Link>
            <Link href={"/docs"}>
              <li>Docs</li>
            </Link>
            <li>
              <AuthButton />
            </li>
          </ul>
        </div>
      </div>
    </header>
  );
};

export default Navbar;
