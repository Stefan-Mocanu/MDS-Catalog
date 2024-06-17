import Plot from "react-plotly.js";
import { parentLoader } from "./Parent";
import { useLoaderData } from "react-router-dom";
import { useState } from "react";

export async function loader({ params }) {
  const data = await parentLoader({ params });
  console.log(data);
  const role = data["rol"];
  let plots = [];

  let plotMedieClasament = null;
  const url =
    "/api/GetMedieClasament?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    role["copil"]["clasa"] +
    "&id_elev=" +
    role["copil"]["id"];
  await fetch(url)
    .then((response) => response.json())
    .then(
      (data) =>
        (plotMedieClasament = (
          <Plot
            key="plotMedieClasament"
            data={data.data}
            layout={data.layout}
          />
        ))
    )
    .catch((error) => console.log(error));
  plots.push({ plot: plotMedieClasament });

  let plotEvolutieElev = null;
  const url2 =
    "/api/GetEvolutieElev?id_scoala=" +
    role["id"] +
    "&id_clasa=" +
    role["copil"]["clasa"] +
    "&id_elev=" +
    role["copil"]["id"];
  await fetch(url2)
    .then((response) => response.json())
    .then(
      (data) =>
        (plotEvolutieElev = (
          <Plot key="plotEvolutieElev" data={data.data} layout={data.layout} />
        ))
    )
    .catch((error) => console.log(error));
  plots.push({ plot: plotEvolutieElev });

  return plots;
}

export default function ParentStatistics() {
  let plots = useLoaderData();
  const [idClasa, setIdClasa] = useState("");
  return (
    <>
      <h2>Parent Statistics</h2>
      {plots.map((plot, i) => {
        return (
          <div key={i}>
            {plot.plot}
            <br />
          </div>
        );
      })}
    </>
  );
}
