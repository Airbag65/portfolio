import React from 'react'
import { NavLink } from 'react-router-dom'

const Navbar = () => {
const linkClass =
        "group flex flex-col items-center";
    const underlineClass =
        "h-1 w-full transition-all duration-200 rounded-t-2xl";

    return (
        <nav className="mx-auto w-full max-w-[20rem] flex justify-between p-1 font-medium sm:pt-3 pb-3 text-black visited:text-black">
            <h2>Anton Norman</h2>
            <div className={linkClass}>
                <NavLink
                    to="/"
                    className={() => ""}
                >
                    Home
                </NavLink>
                <NavLink
                    to="/"
                    end
                    className={({ isActive }) =>
                        `${underlineClass} ${
                            isActive
                                ? "bg-primary-600"
                                : "bg-transparent group-hover:bg-gray-400"
                        }`
                    }
                    tabIndex={-1}
                />
            </div>
            <div className={linkClass}>
                <NavLink
                    to="/about"
                    className={() => ""}
                >
                    About me
                </NavLink>
                <NavLink
                    to="/about"
                    className={({ isActive }) =>
                        `${underlineClass} ${
                            isActive
                                ? "bg-primary-600"
                                : "bg-transparent group-hover:bg-gray-400"
                        }`
                    }
                    tabIndex={-1}
                />
            </div>
        </nav>
    );
}

export default Navbar;
