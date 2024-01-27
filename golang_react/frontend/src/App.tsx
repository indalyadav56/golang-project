import { createBrowserRouter, Outlet, RouterProvider } from 'react-router-dom'

// Pages
import Home from './pages/HomePage'
import Login from './pages/Login'
import ErrorPage from './pages/ErrorPage'

const Layout = () => {
  return (
    <>
      {/* Navbar and other common elements */}

      <h1>NavBar Menu</h1>
      <Outlet />
      {/* Individual Page Components Display Here */}
    </>
  )
}

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        path: '/',
        element: <Home />,
      },
      {
        path: '/login',
        element: <Login />,
      },
    ],
  },
  {
    path: '*',
    element: <ErrorPage />,
  },
])

const App = () => {
  return <RouterProvider router={router} />
}

export default App
