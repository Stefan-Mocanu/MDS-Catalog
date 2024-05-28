import { createContext, useEffect, useState } from "react";
import Login from "./Login.jsx";
import { action as loginAction } from "./Login.jsx";
import Register from "./Register.jsx";
import { action as registerAction } from "./Register.jsx";
import ErrorPage from "./ErrorPage.jsx";
import Profile from "./Profile.jsx";
import { loader as profileLoader } from "./Profile.jsx";
import Layout, { layoutLoader } from "./Layout.jsx";
import Admin from "./Admin.jsx";
import SchoolStatistics from "./SchoolStatistics.jsx";
import { loader as schoolStatisticsLoader } from "./SchoolStatistics.jsx";
import AddAdmin from "./AddAdmin.jsx";
import { action as addAdminAction } from "./AddAdmin.jsx";
import AddUser from "./AddUser.jsx";
import { action as addUserAction } from "./AddUser.jsx";
import Student from "./Student.jsx";
import { loader as studentLoader } from "./Student.jsx";
import ParentAcademicSituation from "./ParentAcademicSituation.jsx";
import Parent from "./Parent.jsx";
// import { loader as parentLoader } from "./Parent.jsx";
import Professor from "./Professor.jsx";
import { professorLoader } from "./Professor.jsx";
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
  redirect,
} from "react-router-dom";

import "./style//App.css";
import Logout from "./Logout.jsx";
import AdminAddUsers from "./AdminAddUsers.jsx";
import { action as AdminAddUsersAction } from "./AdminAddUsers.jsx";
import GetTokens from "./GetTokens.jsx";
import AddAnotherAdmin from "./AddAnotherAdmin.jsx";
import { action as AddAnotherAdminAction } from "./AddAnotherAdmin.jsx";
import StudentAcademicSituation from "./StudentAcademicSituation.jsx";
function App() {
  //request
  const [isSessionActive, setIsSessionActive] = useState(true);
  let userData = null;
  let roluri = null;

  // useEffect(() => {
  const url = "/api/sessionActive";
  fetch(url)
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      if (data !== false) {
        userData = { ...data };
        // if (userData) {

        // }
        // console.log(userData);
        setIsSessionActive(true);
        // console.log(isSessionActive);
      } else {
        userData = null;
        setIsSessionActive(false);
      }
    })
    .catch((error) => console.error("Error:", error));
  // }, []);

  // console.log(roluri);
  // const RoleContext = createContext(null);

  const router = createBrowserRouter([
    {
      path: "/login",
      element: isSessionActive ? <Navigate to="/" /> : <Login />,
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
        <Layout roluri={roluri} />
      ) : (
        <Navigate to="/login" />
      ),
      errorElement: <ErrorPage />,
      loader: layoutLoader,
      children: [
        {
          element: <Logout />,
          path: "/logout",
          errorElement: <ErrorPage />,
        },
        {
          element: <Profile />,
          index: true,
          loader: profileLoader,
        },
        {
          path: "admin/:roleNumber",
          element: <Admin />,
          errorElement: <ErrorPage />,
          // loader: adminLoader,
          children: [
            {
              index: true,
              element: <AdminAddUsers />,
              errorElement: <ErrorPage />,
              action: AdminAddUsersAction,
              // loader: adminLoader,
            },
            {
              path: "gettokens",
              element: <GetTokens />,
              errorElement: <ErrorPage />,
              // loader: adminLoader,
            },
            {
              path: "addanotheradmin",
              element: <AddAnotherAdmin />,
              errorElement: <ErrorPage />,
              action: AddAnotherAdminAction,
            },
            {
              path: "schoolstatistics",
              element: <SchoolStatistics />,
              errorElement: <ErrorPage />,
              loader: schoolStatisticsLoader
            },
          ],
        },
        {
          path: "student/:roleNumber",
          element: <Student />,
          errorElement: <ErrorPage />,
          // loader: studentLoader,
          children: [
            {
              element: <StudentAcademicSituation />,
              index: true,
              errorElement: <ErrorPage />,
            },
          ],
        },
        {
          path: "parent/:roleNumber",
          element: <Parent />,
          errorElement: <ErrorPage />,
          // loader: parentLoader,
          children: [
            {
              element: <ParentAcademicSituation />,
              index: true,
              errorElement: <ErrorPage />,
            },
          ],
        },
        {
          path: "professor/:roleNumber",
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
              action: addAdminAction,
            },
            {
              element: <AddUser />,
              path: "adduser",
              errorElement: <ErrorPage />,
              action: addUserAction,
            },
          ],
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
}

export default App;
