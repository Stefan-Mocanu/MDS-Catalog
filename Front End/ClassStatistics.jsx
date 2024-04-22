import { useLoaderData } from "react-router-dom";

export async function loader({ params }) {
  const classStats = [{ className: "0", professor: "unknown", mean: 5.01 }];
  if (params.idProfessor == 1 && params.idClass == 1)
    classStats.push({
      className: "XII",
      professor: "Paul Ciobanu",
      mean: 9.44,
    });
  return classStats;
}

export default function ClassStatistics() {
  const classStats = useLoaderData();
  return (
    <>
      <h2>
        Statistics for class {classStats[1]["className"]}{" "}
        {classStats[1]["professor"]}: mean = {classStats[1]["mean"]}
      </h2>
    </>
  );
}
