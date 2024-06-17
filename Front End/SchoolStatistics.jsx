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

  let sunBurstIncadrare = null;
  const url5 = "/api/GetSunBurstIncadrare?id_scoala=" + role["id"];
  await fetch(url5)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      sunBurstIncadrare = <Plot data={data.data} layout={data.layout} />;
    });

  let scatterMediiAbsente = null;
  const url6 = "/api/GetScatterMediiAbsente?id_scoala=" + role["id"];
  await fetch(url6)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      scatterMediiAbsente = <Plot data={data.data} layout={data.layout} />;
    });

  let funnelMedii = null;
  const url7 = "/api/GetFunnelMedii?id_scoala=" + role["id"];
  await fetch(url7)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      funnelMedii = <Plot data={data.data} layout={data.layout} />;
    });

  let heatMapMedii = null;
  const url8 = "/api/GetHeatmapMediiActiv?id_scoala=" + role["id"];
  await fetch(url8)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      heatMapMedii = <Plot data={data.data} layout={data.layout} />;
    });

  let heatMapIncadrare = null;
  const url9 = "/api/GetHeatmapIncadrare?id_scoala=" + role["id"];
  await fetch(url9)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      heatMapIncadrare = <Plot data={data.data} layout={data.layout} />;
    });

  let parralelCatFeedback = null;
  const url10 = "/api/GetParralelCatFeedback?id_scoala=" + role["id"];
  await fetch(url10)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      parralelCatFeedback = <Plot data={data.data} layout={data.layout} />;
    });

  let feedbackChart = null;
  const url11 = "/api/GetFeedbackChart?id_scoala=" + role["id"];
  await fetch(url11)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      feedbackChart = <Plot data={data.data} layout={data.layout} />;
    });

    let procPozitiv = null;
    const url12 = "/api/GetProcPozitiv?id_scoala=" + role["id"];
  await fetch(url12)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      procPozitiv = <Plot data={data.data} layout={data.layout} />;
    });

    let heatMapMediiLuni = null;
    const url13 = "/api/HM_MediiLuniA?id_scoala=" + role["id"];
  await fetch(url13)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      heatMapMediiLuni = <Plot data={data.data} layout={data.layout} />;
    });


    let eleviPromov = null;
    const url14 = "/api/EleviPromov?id_scoala=" + role["id"];
  await fetch(url14)
    .then((response) => response.json())
    .then((data) => {
      // console.log(data)
      eleviPromov = <Plot data={data.data} layout={data.layout} />;
    })
    .catch((error) => console.log(error));


  let plots = {
    distEtniiPlot: distEtniiPlot,
    devStdNote: devStdNote,
    repMediiEtnii: repMediiEtnii,
    repMediiGen: repMediiGen,
    sunBurstIncadrare: sunBurstIncadrare,
    scatterMediiAbsente: scatterMediiAbsente,
    funnelMedii: funnelMedii,
    heatMapMedii: heatMapMedii,
    heatMapIncadrare: heatMapIncadrare,
    parralelCatFeedback: parralelCatFeedback,
    feedbackChart: feedbackChart,
    procPozitiv: procPozitiv,
    heatMapMediiLuni: heatMapMediiLuni,
    eleviPromov:eleviPromov
  };
  return plots;
}

export default function SchoolStatistics() {
  const plots = useLoaderData();
  const distEtniiPlot = plots["distEtniiPlot"];
  const devStdNote = plots["devStdNote"];
  const repMediiEtnii = plots["repMediiEtnii"];
  const repMediiGen = plots["repMediiGen"];
  const sunBurstIncadrare = plots["sunBurstIncadrare"];
  const scatterMediiAbsente = plots["scatterMediiAbsente"];
  const funnelMedii = plots.funnelMedii;
  const heatMapMedii = plots.heatMapMedii;
  const heatMapIncadrare = plots.heatMapIncadrare;
  const parralelCatFeedback = plots.parralelCatFeedback;
  const feedbackChart = plots.feedbackChart;
  const procPozitiv = plots.procPozitiv;
const heatMapMediiLuni = plots.heatMapMediiLuni;
const eleviPromov = plots.eleviPromov;

  return (
    <>
      <h2>SchoolStatistics</h2>
      {sunBurstIncadrare}
      <br />
      <br />
      <br />
      <br />
      <h4>Etnies distribution in school</h4>
      {distEtniiPlot}
      <br />
      <br />
      <h4>Box Plot: Repartition of the average grades by etnies</h4>
      {repMediiEtnii}
      <br />
      <br />
      <h4>Box Plot: Repartition of the average grades by gender</h4>
      {repMediiGen}
      <br />
      <br />
      <h4>Standard deviation of the grades in school</h4>
      {devStdNote}
      <br />
      <br />
      {scatterMediiAbsente}
      <br />
      <br />
      {funnelMedii}
      <br />
      <br />
      {heatMapMedii}
      <br />
      <br />
      {heatMapIncadrare}
      <br />
      <br />
      {parralelCatFeedback}
      <br />
      <br />
      {feedbackChart}
      <br />
      <br />
      {procPozitiv}
      <br />
      <br />
      {heatMapMediiLuni}
      <br />
      <br />
      {eleviPromov}
    </>
  );
}
