import {
  Link,
  NavLink,
  Outlet,
  useLoaderData,
  useOutletContext,
  useParams,
} from "react-router-dom";

function getRol(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Administrator") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

export default function Admin() {
  // const data = useLoaderData();
  const roluri = useOutletContext();
  const roleNumber = useParams()["roleNumber"];
  const data = getRol(roluri, roleNumber);
  let context = data;
  console.log(data);
  let pathToThisPage = "/admin/" + roleNumber;
  return (
    <>
      <div id="useroptions">
        <NavLink
          to={pathToThisPage}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Add school information</button>
        </NavLink>
        <NavLink
          to={"gettokens"}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Get tokens for the users of the school</button>
        </NavLink>
        <NavLink
          to={"schoolstatistics"}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>View school statistics</button>
        </NavLink>
        <NavLink
          to={"addanotheradmin"}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Add another admin</button>
        </NavLink>
      </div>
      <div id="content">
        <h4>Admin {data["firstName"]}</h4>
        <Outlet context={context} />
      </div>
    </>
  );
}
