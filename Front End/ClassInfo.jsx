import { useLoaderData } from "react-router-dom";

export async function loader({ params }) {
  const classInfo = [{ className: "0", professor: "unknown", nrOfStuds: 30 }];
  if (params.idProfessor == 1 && params.idClass == 1)
    classInfo.push({
      className: "XII",
      professor: "Paul Ciobanu",
      nrOfStuds: 15,
    });
  return classInfo;
}

export default function ClassInfo() {
  const classInfo = useLoaderData();
  return (
    <>
      <h2>
        Information about class {classInfo[1]["className"]}{" "}
        {classInfo[1]["professor"]}
      </h2>
    </>
  );
}
