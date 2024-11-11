import React from 'react';
import { Link } from 'react-router-dom';

// Import your logo (replace with your logo path)
import logo from '../../assets/logo.png';

const Navbar: React.FC = () => {
  return (
    <nav className="bg-blue-500 p-4 text-white">
      <div className="container mx-auto flex justify-between items-center">
        {/* Left Links */}
        <div className="flex justify-start space-x-4">
          <Link to="/" className="font-semibold text-xl">
            Ticketing System
          </Link>
        </div>

        {/* Center Logo */}
        <div className="flex justify-center flex-grow">
          <img src={logo} alt="Logo" className="h-12" />
        </div>

        {/* Right Links */}
        <div className="flex justify-end space-x-4">
          <Link to="/ticket-list" className="hover:underline">
            Tickets
          </Link>
          <Link to="/create-ticket" className="hover:underline">
            Create Ticket
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
