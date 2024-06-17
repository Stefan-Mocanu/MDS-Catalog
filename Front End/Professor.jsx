import {
  NavLink,
  Outlet,
  useLoaderData,
  useOutletContext,
  useParams,
} from "react-router-dom";
import { layoutLoader } from "./Layout";

function getRol(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Profesor") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

async function getClasses(role) {
  const url = "/api/claseProfesor";
  let formData = new FormData();
  let classes;
  formData.append("id_scoala", role["id"]);
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => response.json())
    .then((data) => (classes = data["data"]))
    .catch((error) => console.log(error));
  return classes;
}

export async function professorLoader({ params }) {
  const roluri = await layoutLoader();
  const role = getRol(roluri, params.roleNumber);
  const classes = await getClasses(role);
  const data = { role: role, classes: classes };
  return data;
}

export default function Professor() {
  // const data = useLoaderData();
  const roleNumber = useParams()["roleNumber"];
  // console.log(data);

  let pathToThisPage = "/professor/" + roleNumber;
  return (
    <>
      <div id="useroptions">
        <NavLink
          to={pathToThisPage}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Classes</button>
        </NavLink>
        <NavLink
          to={"professorfeedback"}
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Feedback</button>
        </NavLink>
        {/* <NavLink
          to={"professorstatistics"}
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Statistics</button>
        </NavLink> */}
      </div>
      <div id="content">
        <h2></h2>
        <Outlet />
      </div>
    </>
  );
}
