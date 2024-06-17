import { useEffect, useState } from "react";
import {
  useLoaderData,
  Outlet,
  useOutletContext,
  useParams,
  NavLink,
} from "react-router-dom";
import { layoutLoader } from "./Layout";

function getRole(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Elev") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

async function getClase(id_scoala) {
  let clase;
  const url = "/api/getClase?id_scoala=" + id_scoala;
  await fetch(url)
    .then((response) => response.json())
    .then((data) => {
      clase = data;
    })
    .catch((error) => console.log(error));
  return clase;
}

async function getCatalog(id_scoala, id_clasa) {
  let catalog;
  const url =
    "/api/viewCatalogElev?id_scoala=" + id_scoala + "&id_clasa=" + id_clasa;
  await fetch(url)
    .then((response) => response.json())
    .then((data) => {
      catalog = { id_clasa: id_clasa, catalog: data };
    })
    .catch((error) => {
      console.log(error);
    });
  return catalog;
}

export async function studentLoader({ params }) {
  const roluri = await layoutLoader();
  const role = getRole(roluri, params.roleNumber);
  const clase = await getClase(role["id"]);
  let cataloage = [];
  for (let i = 0; i < clase.length; i++) {
    const catalog = await getCatalog(role["id"], clase[i]);
    cataloage.push(catalog);
  }
  return { rol: role, cataloage: cataloage };
}

export default function Student() {
  const data = useLoaderData();
  console.log(data);
  const roleNumber = useParams()["roleNumber"];
  const role = data["rol"];
  const cataloage = data["cataloage"];
  const context = { rol: role, cataloage: cataloage };
  let pathToThisPage = "/student/" + roleNumber;
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
        <NavLink 
          to={"feedbackforprofessors"}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Feedback for professors</button></NavLink>
      </div>
      <div id="content">
        <Outlet context={context} />
      </div>
    </>
  );
}
