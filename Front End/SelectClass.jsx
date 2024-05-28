import { useLoaderData } from "react-router-dom";
import { professorLoader } from "./Professor";

export async function loader(props) {
  return professorLoader(props);
}

export default function SelectClass() {
  const data = useLoaderData();
  const role = data["role"];
  const classes = data["classes"];
  
  return (
    <>
      <h2>Select Class</h2>
      
    </>
  );
}
