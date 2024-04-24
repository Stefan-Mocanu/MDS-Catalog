import { useState } from "react";
import { Navigate, redirect } from "react-router-dom";

export default function Logout() {
  const url = "/api/logout";
  fetch(url, {
    method: "POST",
  });
  //   const [ceva, setCeva] = useState("ceva");
  //   if (ceva === "ceva") setCeva({ ceva: "ceva" });
  //   else if (ceva == { ceva: "ceva" }) setCeva({ nustiu: "nustiu" });
  location.replace("/");
  redirect("/login");
  return <Navigate to="/login" />;
}
