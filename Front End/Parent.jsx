import { Outlet, useLoaderData } from "react-router-dom";

export async function loader({ params }) {
  const parentData = { firstName: "NoFirstName", lastName: "NoLastName" };
  if (params.idParent == 1) {
    parentData["firstName"] = "Paul";
    parentData["lastName"] = "Ciobanu";
  }
  //de preluat si date despre copilul aferent rolului
  return parentData;
}

export default function Parent() {
  const parentData = useLoaderData();
  return (
    <>
      <div id="useroptions">
        <button>Academic situation</button>
        <button>Feedback from professors</button>
        <button>Achievements and badges</button>
        <button>Child's progress and graphics</button>
        <button>Discuss with professors</button>
      </div>
      
      <div id="content">
        <h2>
          Parent {parentData["firstName"]} {parentData["lastName"]}
        </h2>
        <Outlet />
      </div>
    </>
  );
}
