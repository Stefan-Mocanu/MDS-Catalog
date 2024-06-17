import { useEffect, useState } from "react";
import { NavLink, Outlet, useLoaderData, useParams } from "react-router-dom";
import { layoutLoader } from "./Layout";

function getRole(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Parinte") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

async function getCatalog(id_scoala, id_clasa, id_elev) {
  let catalog;
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
      catalog = data;
    })
    .catch((error) => {
      console.log(error);
    });
  return catalog;
}

export async function parentLoader({ params }) {
  const roluri = await layoutLoader();
  const role = getRole(roluri, params.roleNumber);
  const catalog = await getCatalog(
    role["id"],
    role["copil"]["clasa"],
    role["copil"]["id"]
  );
  return { rol: role, catalog: catalog };
}

export default function Parent() {
  const data = useLoaderData();
  const roleNumber = useParams()["roleNumber"];
  const catalog = data["catalog"];

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

        <NavLink
          to={"parentstatistics"}
          end
          className={({ isActive }) => (isActive ? "selectedbutton" : "")}
        >
          <button>Child's progress and graphics</button>
        </NavLink>
      </div>

      <div id="content">
        <h2></h2>
        {catalog && <Outlet context={data} />}
      </div>
    </>
  );
}
