import { Outlet } from "react-router-dom";
import NavBar from "./NavBar";
import "./style/Layout.css";
export default function Layout({ userData, setUserData }) {
  return (
    <>
      <div id="layout">
        <NavBar />
        <div id="rolecontent">
            <Outlet />
        </div>
      </div>
    </>
  );
}
