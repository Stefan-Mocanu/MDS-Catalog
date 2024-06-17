import { Form, useActionData, useLoaderData } from "react-router-dom";
import { professorLoader } from "./Professor";
import Plot from "react-plotly.js";

export async function loader({ params }) {
  const data = await professorLoader({ params });
  const role = data["role"]
  let plots = [];

  let mediiClaseProf = null;
  const url = "/api/MediiClase?id_scoala=" + role["id"];
  await fetch(url)
    .then((response) => response.json())
    .then(
      (data) =>
        {mediiClaseProf = <Plot data={data.data} layout={data.layout} />}
    )
    .catch((error) => console.log(error));

  plots.push(mediiClaseProf);
  return plots;
}

export async function action({ request }) {
  let ceva = "Salut";
  console.log(ceva);
  return ceva;
}

export default function ProfessorStatistics() {
  console.log("in component");
  let plots = useLoaderData();
  return (
    <>
      <h2>ProfessorStatistics</h2>
      {plots.map((plot, i) => {
        return <div key={i}>{plot}</div>;
      })}
    </>
  );
}
