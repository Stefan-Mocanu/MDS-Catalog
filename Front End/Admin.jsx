import { Link, Outlet, useLoaderData } from "react-router-dom";

export async function loader({ params }) {
  var adminData = { firstName: "firstName" };
  // const url = "/api/sessionActive";
  // await fetch(url)
  //   .then((response) => response.json())
  //   .then((data) => {
  //     adminData = data;
  //   })
  //   .catch((error) => console.error("Error:", error));
  // if (adminData.id != params.idAdmin) throw new Error("Invalid parameter!");
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
