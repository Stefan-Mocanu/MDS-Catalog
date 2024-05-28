import { useState } from "react";
import { Link, NavLink } from "react-router-dom";
function linkFromRole(rol, i) {
  let tip;
  switch (rol) {
    case "Administrator":
      tip = "admin";
      break;
    case "Elev":
      tip = "student";
      break;
    case "Parinte":
      tip = "parent";
      break;
    case "Profesor":
      tip = "professor";
      break;
  }
  return "/" + tip + "/" + i;
}
export default function NavBar({ roluri }) {
  console.log("From navbar:");
  console.log(roluri);
  let buttons = [];
  for (let i=0; i< roluri.length; i++) {
    let elev = roluri[i]["copil"];
    let numeElev = "";
    let clasa = ""
    if (elev != null) {
      numeElev = elev["nume"] + " " + elev["prenume"];
      clasa = elev["clasa"];
    }

    buttons.push(
      <NavLink
        className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        key={i}
        to={linkFromRole(roluri[i]["rol"], (i+1))}
      >
        <button>
          {roluri[i]["rol"]} <br />
          {numeElev + " " + clasa + " "}  
          {roluri[i]["scoala"]}
        </button>
      </NavLink>
    );
  }
  return (
    <div id="nav_bar">
      <div id="generaloptions">
        <Link to="/logout">
          <button>Logout</button>
        </Link>
        <Link to="/">
          <button>Profile</button>
        </Link>
        <Link to="/addrole">
          <button>Add role</button>
        </Link>
      </div>
      <div>
        <h3>Roles:</h3>
      </div>
      <div id="roles">{buttons}</div>
    </div>
  );
}
