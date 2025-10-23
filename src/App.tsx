import React from "react"
import { Routes, Route } from "react-router-dom"
import Root from "./Root.tsx"
import HomePage from "./pages/HomePage.tsx"
import AboutMe from "./pages/AboutMe.tsx"
// import PageBrowse from "./pages/PageBrowse"
// import PageMyLibrary from "./pages/PageMyLibrary"
// import PageViewer3D from "./pages/PageViewer3D"


export default function App() {
  return (
    // If using lazy() above, wrap <Routes> in <React.Suspense fallback={...}>
    <Routes>
      <Route path="/" element={<Root />}>
        <Route index element={<HomePage />} />

        <Route path="/about" element={<AboutMe />} />
      </Route>
    </Routes>
  )
}
// <Route path="browse" element={<PageBrowse />} />
// <Route path="my-library" element={<PageMyLibrary />} />
// <Route path="viewer" element={<PageViewer3D />} />
