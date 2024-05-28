import { useLoaderData } from "react-router-dom";
import { layoutLoader } from "./Layout";
import Plot from "react-plotly.js";

function getRol(roluri, roleNumber) {
  const rol = roluri[roleNumber - 1];
  if (rol["rol"] !== "Administrator") {
    throw new Response("Not Found", { status: 404 });
  }
  return rol;
}

export async function loader({ params }) {
  const roluri = await layoutLoader();
  const role = getRol(roluri, params.roleNumber);
  const url = "/api/getDistEtnii?id_scoala=" + role["id"];
  let distEtniiPlot = null;
  await fetch(url)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      distEtniiPlot = <Plot data={data.data} layout={data.layout} />;
    });

  let devStdNote = null;
  const url2 = "/api/getDevStdNote?id_scoala=" + role["id"];
  await fetch(url2)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      devStdNote = <Plot data={data.data} layout={data.layout} />;
    });

  let repMediiEtnii = null;
  const url3 = "/api/getRepMediiEtnii?id_scoala=" + role["id"];
  await fetch(url3)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      repMediiEtnii = <Plot data={data.data} layout={data.layout} />;
    });

  let repMediiGen = null;
  const url4 = "/api/getRepMediiGen?id_scoala=" + role["id"];
  await fetch(url4)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      repMediiGen = <Plot data={data.data} layout={data.layout} />;
    });

  let plots = {
    distEtniiPlot: distEtniiPlot,
    devStdNote: devStdNote,
    repMediiEtnii: repMediiEtnii,
    repMediiGen: repMediiGen,
  };
  return plots;
}

export default function SchoolStatistics() {
  const plots = useLoaderData();
  const distEtniiPlot = plots["distEtniiPlot"];
  const devStdNote = plots["devStdNote"];
  const repMediiEtnii = plots["repMediiEtnii"];
  const repMediiGen = plots["repMediiGen"];
  return (
    <>
      <h2>SchoolStatistics</h2>
      <h4>Etnies distribution in school</h4>
      {distEtniiPlot}
      <br />
      <br />
      <h4>Box Plot: Repartitia mediilor pe etnii</h4>
      {repMediiEtnii}
      <br />
      <br />
      {repMediiGen}
      <br />
      <br />
      <h4>Standard deviation of the grades in school</h4>
      {devStdNote}
    </>
  );
}
