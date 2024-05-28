import { useEffect, useState } from "react";
import {
  useLoaderData,
  Outlet,
  useOutletContext,
  useParams,
  NavLink,
} from "react-router-dom";

export async function loader({ params }) {
  const studentData = { firstName: "NoLastName", lastName: "NoFirstName" };
  if (params.idStudent == 1) {
    studentData["firstName"] = "Paul";
    studentData["lastName"] = "Ciobanu";
  }
  return studentData;
}

function getRole(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Elev") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

export default function Student() {
  // const studentData = useLoaderData();
  const roluri = useOutletContext();
  const roleNumber = useParams()["roleNumber"];
  const role = getRole(roluri, roleNumber);
  const [clase, setClase] = useState([]);
  const [cataloage, setCataloage] = useState([]);
  useEffect(() => {
    async function getClase(id_scoala) {
      const url = "/api/getClase?id_scoala=" + id_scoala;
      await fetch(url)
        .then((response) => response.json())
        .then((response) => {
          setClase(response);
          console.log(response);
        })
        .catch((error) => console.log(error));
      return 0;
    }
    getClase(role["id"]);
  }, [role]);

  useEffect(() => {
    async function getCatalog(id_scoala, id_clasa, i) {
      if(i === 0)
        setCataloage([]);
      const url =
        "/api/viewCatalogElev?id_scoala=" + id_scoala + "&id_clasa=" + id_clasa;
      await fetch(url)
        .then((response) => response.json())
        .then((response) => {
          let obj = { id_clasa: id_clasa, catalog: response };
          console.log(cataloage);
          setCataloage([...cataloage, obj]);
          console.log(cataloage);
        })
        .catch((error) => {
          console.log(error);
        });
      return 0;
    }
    setCataloage([]);
    // location.reload();
    clase.forEach((id_clasa, i) => {
      getCatalog(role["id"], id_clasa, i);
    });
  }, [clase]);

  console.log(clase);
  console.log(cataloage);
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
        <button>Feedback from professors</button>
        <button>Achievements and badges</button>
      </div>
      <div id="content">
        <h2>
          {/* Student {studentData["firstName"]} {studentData["lastName"]} */}
        </h2>
        <Outlet context={context} />
      </div>
    </>
  );
}
