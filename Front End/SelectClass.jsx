import { useLoaderData } from "react-router-dom";

export async function loader({ params }) {
  const classes = [{ className: "0", numStudents: 30 }];
  if (params["idProfessor"] == 1)
    classes.push({ className: "XII", numStudents: 15 });
  return classes;
}

export default function SelectClass() {
  const classes = useLoaderData();
  return (
    <>
      <h2>Select Class</h2>
      <div>
        Class {classes[1]["className"]} has {classes[1]["numStudents"]}{" "}
        students.
      </div>
    </>
  );
}
