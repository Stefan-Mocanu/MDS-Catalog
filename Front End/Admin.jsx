import { Link, Outlet, useLoaderData } from "react-router-dom";

export async function loader({ params }) {
  let adminData = { firstName: "NoFirstName", lastName: "NoLast" };
  if (params.idAdmin == 1) {
    adminData["firstName"] = "Paul";
    adminData["lastName"] = "Ciobanu";
  }
  return adminData;
}

export default function Admin() {
  const data = useLoaderData();
  console.log(data);
  return (
    <>
      <div id="useroptions">
        <button>Add users for school</button>
        <button>Edit school info??</button>
        <button>View school statistics</button>
        <button>Optiune 2 oarecare</button>
      </div>
      <div id="content">
        <h4>Admin {data["firstName"]}</h4>
        <Outlet />
      </div>
    </>
  );
}
