import { Outlet, useLoaderData } from "react-router-dom";
import NavBar from "./NavBar";
import "./style/Layout.css";
export async function layoutLoader() {
  let roluri = [];
  const url = "/api/getRoluri";
  await fetch(url)
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      roluri = { ...data };
    })
    .catch((error) => console.error("Error:", error));
  return roluri;
}
export default function Layout() {
  const roluri = useLoaderData();
  console.log("From layout:");
  console.log(roluri);
  return (
    <>
      <div id="layout">
        <NavBar roluri={roluri} />
        <div id="rolecontent">
          <Outlet />
        </div>
      </div>
    </>
  );
}
