"use client";
import { signIn, signOut, useSession } from "next-auth/react";
import { usePathname } from "next/navigation";

import React from "react";

const AuthButton = () => {
  const { data: session } = useSession();
  //   const pathname = usePathname();

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
    <div>
      <AuthButton />
    </div>
  );
};

export default Navbar;
