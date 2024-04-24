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
  for (let key in roluri) {
    buttons.push(
      <NavLink key={key} to={linkFromRole(roluri[key]["rol"], key)}>
        <button>
          {roluri[key]["rol"]} {roluri[key]["scoala"]}
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
        <button>Profile</button>
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
