// export async function action({ request }) {
//     let formData = await request.formData();
//     let intent = formData.get("intent");

import { useLoaderData, useOutletContext } from "react-router-dom";

//     if (intent === "getprofessors") {

//     }

//     if (intent === "getstudents") {

//     }
//     else
//         throw new Response("Not Found", { status: 404 });

// }

export default function GetTokens() {
  const data = useOutletContext();
  const urlprof = "/api/csvProfesor?id_scoala=" + data["id"];
  const urlstud = "/api/csvElev?id_scoala=" + data["id"];

  return (
    <>
      {/* <form method="get"> */}
      <label>
        Get tokens for professors:
        <br />
        <button name="intent" value="getprofessors">
          <a href={urlprof}>Get tokens</a>
        </button>
      </label>
      <br />
      {/* </form> */}
      {/* <form method="get"> */}
      <label>
        Get tokens for students and parents:
        <br />
        <button name="intent" value="getstudents">
          <a href={urlstud}>Get tokens</a>
        </button>
      </label>
      {/* </form> */}
    </>
  );
}
