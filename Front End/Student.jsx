import { useLoaderData, Outlet } from "react-router-dom";

export async function loader({ params }) {
  const studentData = { firstName: "NoLastName", lastName: "NoFirstName" };
  if (params.idStudent == 1) {
    studentData["firstName"] = "Paul";
    studentData["lastName"] = "Ciobanu";
  }
  return studentData;
}
export default function Student() {
  const studentData = useLoaderData();

  return (
    <>
      <div id="useroptions">
        <button>Academic situation</button>
        <button>Feedback from professors</button>
        <button>Achievements and badges</button>
      </div>
      <div id="content">
        <h2>
          Student {studentData["firstName"]} {studentData["lastName"]}
        </h2>
        <Outlet />
      </div>
    </>
  );
}
