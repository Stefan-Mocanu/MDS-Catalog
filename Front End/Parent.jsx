import { useEffect, useState } from "react";
import {
  NavLink,
  Outlet,
  useLoaderData,
  useOutletContext,
  useParams,
} from "react-router-dom";

function getRole(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Parinte") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

export default function Parent() {
  const roluri = useOutletContext();
  const roleNumber = useParams()["roleNumber"];
  const role = getRole(roluri, roleNumber);
  const [catalog, setCatalog] = useState(null);
  const [context, setContext] = useState(null);
  useEffect(() => {
    async function getCatalog(id_scoala, id_clasa, id_elev) {
      let url =
        "/api/viewCatalogParinte?id_scoala=" +
        id_scoala +
        "&id_clasa=" +
        id_clasa +
        "&id_elev=" +
        id_elev;
      await fetch(url)
        .then((response) => response.json())
        .then((data) => {
          setCatalog(data);
          setContext({ rol: role, catalog: data });
        })
        .catch((error) => {
          console.log(error);
        });
      return catalog;
    }
    getCatalog(role["id"], role["copil"]["clasa"], role["copil"]["id"]);

  }, [role]);
  console.log(catalog);

  let pathToThisPage = "/parent/" + roleNumber;
  return (
    <>
      <div id="useroptions">
        <NavLink
          to={pathToThisPage}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Student's academic situation</button>
        </NavLink>
        <button>Feedback from professors</button>
        <button>Achievements and badges</button>
        <button>Child's progress and graphics</button>
        <button>Discuss with professors</button>
      </div>

      <div id="content">
        <h2></h2>
        {catalog && <Outlet context={context} />}
      </div>
    </>
  );
}
