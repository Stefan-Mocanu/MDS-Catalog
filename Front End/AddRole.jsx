import AddAdmin from "./AddAdmin";
import AddUser from "./AddUser";
import { useState } from "react";
import { Form, NavLink, Outlet } from "react-router-dom";

export async function action({ request }) {
  async function fun(request) {
    setTimeout(async () => {
      let formData = await request.formData();
      console.log(formData.get("name"));
    }, 2000);
    return true;
  }
  return fun(request);
}

function handleRadio(e, n, setOpt) {
  // e.preventDefault();
  setOpt(n);
  // console.log("Hello from radio handler");
}
export default function AddRole() {
  const [opt, setOpt] = useState(0);

  let component;
  if (opt == 0) component = <div></div>;
  else if (opt == 1) component = <AddUser />;
  else component = <AddAdmin />;

  return (
    <>
      <h2>Add role</h2>
        <p>Choose the type of user:</p>
        <NavLink
          to="/addrole/adduser"
          className={({ isActive, isPending }) =>
            isActive ? "selectedbutton" : ""
          }
        >
          <button type="button">Add role using received token</button>
        </NavLink>
        <br></br>
        <NavLink
          to="/addrole/addadmin"
          className={({ isActive, isPending }) =>
            isActive ? "selectedbutton" : ""
          }
        >
          <button type="button">Add admin role and register a school</button>
        </NavLink>
        <div>
          <Outlet />
        </div>
    </>
  );
}
