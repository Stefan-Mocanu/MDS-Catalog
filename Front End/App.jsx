import { useState } from "react";
import Login from "./Login.jsx";
import { action as loginAction } from "./Login.jsx";
import Register from "./Register.jsx";
import { action as registerAction } from "./Register.jsx";
import ErrorPage from "./ErrorPage.jsx";
import Menu from "./Menu.jsx";
import Layout from "./Layout.jsx";
import Admin from "./Admin.jsx";
import SchoolInfo from "./SchoolInfo.jsx";
import { loader as adminLoader } from "./Admin.jsx";
import { action as roleAction } from "./AddRole.jsx";
import Student from "./Student.jsx";
import { loader as studentLoader } from "./Student.jsx";
import StudentAcademicSituation from "./StudentAcademicSituation.jsx";
import Parent from "./Parent.jsx";
import { loader as parentLoader } from "./Parent.jsx";
import Professor from "./Professor.jsx";
import { loader as professorLoader } from "./Professor.jsx";
import SelectClass from "./SelectClass.jsx";
import { loader as selectClassLoader } from "./SelectClass.jsx";
import ClassStatistics from "./ClassStatistics.jsx";
import { loader as classStatisticsLoader } from "./ClassStatistics.jsx";
import ClassInfo from "./ClassInfo.jsx";
import { loader as classInfoLoader } from "./ClassInfo.jsx";
import AddRole from "./AddRole.jsx";
import {
  Routes,
  Route,
  createBrowserRouter,
  Navigate,
  useLocation,
  RouterProvider,
  Outlet,
} from "react-router-dom";

import "./style//App.css";
import AddAdmin from "./AddAdmin.jsx";
import AddUser from "./AddUser.jsx";

function App() {
  //request
  const [isSessionActive, setIsSessionActive] = useState(true);
  const [userData, setUserData] = useState({ nume: "Default" });
  const nume = "ceva";
  console.log(userData);
  /////??
  // if (isSessionActive) {
  //   //request
  //   setUserData({ name: "Gica", varsta: 9, tip: "elev" });
  //   return <Home userData={userData} />;
  // }
  /////??
  // loaders?
  // const location = useLocation();

  const router = createBrowserRouter([
    {
      path: "/login",
      element: <Login />,
      errorElement: <ErrorPage />,
      action: loginAction,
    },
    {
      path: "/register",
      element: <Register />,
      errorElement: <ErrorPage />,
      action: registerAction,
    },
    {
      path: "/",
      element: isSessionActive ? (
        <Layout userData={userData} setUserData={setUserData} />
      ) : (
        <Navigate to="/login" />
      ),
      errorElement: <ErrorPage />,
      children: [
        {
          element: <Menu userData={userData} setUserData={setUserData} />,
          index: true,
        },
        {
          path: "admin/:idAdmin",
          element: <Admin />,
          errorElement: <ErrorPage />,
          loader: adminLoader,
          children: [
            {
              element: <SchoolInfo />,
              index: true,
              errorElement: <ErrorPage />,
            },
          ],
        },
        {
          path: "student/:idStudent",
          element: <Student />,
          errorElement: <ErrorPage />,
          loader: studentLoader,
          children: [
            {
              element: <StudentAcademicSituation />,
              index: true,
              errorElement: <ErrorPage />,
            },
          ],
        },
        {
          path: "parent/:idParent",
          element: <Parent />,
          errorElement: <ErrorPage />,
          loader: parentLoader,
          children: [
            {
              element: <StudentAcademicSituation />,
              index: true,
              errorElement: <ErrorPage />,
            },
          ],
        },
        {
          path: "professor/:idProfessor",
          element: <Professor />,
          errorElement: <ErrorPage />,
          loader: professorLoader,
          children: [
            {
              // path: "selectclass",
              element: <SelectClass />,
              index: true,
              errorElement: <ErrorPage />,
              loader: selectClassLoader,
            },
            {
              path: "classstatistics/:idClass",
              element: <ClassStatistics />,
              errorElement: <ErrorPage />,
              loader: classStatisticsLoader,
            },
            {
              path: "classinfo/:idClass",
              element: <ClassInfo />,
              errorElement: <ErrorPage />,
              loader: classInfoLoader,
            },
            // {
            //   path: "allstudentsstatistics",
            //   element: <AllStudentsStatistics />,
            //   errorElement: <ErrorPage />,
            //   loader: AllStudentsStatisticsLoader,
            // },
            // {
            //   path: "allstudentsinfo",
            //   element: <AllStudentsInfo />,
            //   errorElement: <ErrorPage />,
            //   loader: AllStudentsInfoLoader,
            // },
          ],
        },
        {
          path: "addrole",
          element: <AddRole />,
          errorElement: <ErrorPage />,
          children: [
            {
              element: <AddAdmin />,
              path: "addadmin",
              errorElement: <ErrorPage />,
              action: roleAction,
            },
            {
              element: <AddUser />,
              path: "adduser",
              errorElement: <ErrorPage />,
              //action????
            },
          ],
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
}

export default App;
