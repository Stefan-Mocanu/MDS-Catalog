import { Outlet, useLoaderData } from "react-router-dom";
export async function loader({ params }) {
  const profData = { firstName: "NoFirstName", lastName: "NoLastName" };
  if (params["idProfessor"] == 1) {
    profData["firstName"] = "Paul";
    profData["lastName"] = "Ciobanu";
  }
  return profData;
}

export default function Professor() {
  const profData = useLoaderData();

  return (
    <>
      <div id="useroptions">
        <button>Classes</button>
        <button>All students</button>
        <button>Statistics & feedback</button>
        <button>Chat</button>
      </div>
      <div id="content">
        <h2>
          Professor {profData["firstName"]} {profData["lastName"]}
        </h2>
        <Outlet />
      </div>
    </>
  );
}
