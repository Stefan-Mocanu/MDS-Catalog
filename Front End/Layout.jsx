import { Outlet, useLoaderData } from "react-router-dom";
import NavBar from "./NavBar";
import "./style/Layout.css";

async function getChildren(id_scoala){
  const url = "/api/getElevi?id_scoala="+ id_scoala;
  let elevi;
  await fetch(url)
    .then((response)=>response.json())
    .then((data)=>{elevi=data})
    .catch((error)=>console.log(error));
  return elevi;
}

export async function parseData(data){
  let roluri = [];
  for(let key in data){
    if(data[key]["rol"] === "Parinte"){
      let elevi = await getChildren(data[key]["id"]);
      for(let elev of elevi){
        let newRole = data[key];
        newRole["copil"] = elev;
        roluri.push(newRole);
      }
    }
    else{
      let newRole = data[key];
      newRole["copil"] = null;
      roluri.push(newRole);
    }
  }
  return roluri;
}


export async function layoutLoader() {
  let roluri = [];
  const url = "/api/getRoluri";
  await fetch(url)
    .then((response) => response.json())
    .then(async (data) => {
      console.log(data);
      roluri = await parseData(data);
    })
    .catch((error) => console.error("Error:", error));
  return roluri;
}
export default function Layout() {
  const roluri = useLoaderData();
  console.log("From layout:");
  console.log(roluri);
  let context = roluri;
  return (
    <>
      <div id="layout">
        <NavBar roluri={roluri} />
        <div id="rolecontent">
          <Outlet context={context}/>
        </div>
      </div>
    </>
  );
}
