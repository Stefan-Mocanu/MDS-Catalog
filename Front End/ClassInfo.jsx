import { useLoaderData } from "react-router-dom";
import { professorLoader } from "./Professor";


async function getStudents(id_clasa, id_scoala){
  let formData = new FormData();
  formData.append("id_scoala", id_scoala);
  formData.append("id_clasa", id_clasa);
  const url = "/api/eleviClasa";
  let students = [];
  await fetch(url, {
    method: "POST",
    body: formData
  })
  .then(response => response.json())
  .then(data => students = data)
  .catch(error => console.log(error))

  return students;
}

export async function loader({ params }) {
//de  vazut de ce nu merge
  const data = professorLoader({params});
  console.log(data);
  console.log(data);
  const classes = data["classes"];
  const role = data["role"];
  let thisClass = null;
  for(let obj in classes)
    if(obj["clasa"] === params.idClass)
      thisClass = obj;
  if(thisClass == null)
    throw new Response("Not Found", { status: 404 });
  const students = await getStudents(thisClass["clasa"], role["id"]);
  return {thisClass: thisClass, role:role, students:students};
}

export default function ClassInfo() {
  const data = useLoaderData();
  console.log(data);
  return (
    <>
      <h2>
        ceva
      </h2>
    </>
  );
}
