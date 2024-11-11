import React from 'react';
import { Link } from 'react-router-dom';

const Navbar: React.FC = () => {
  return (
    <nav className="bg-blue-500 p-4 text-white">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="font-semibold text-xl">
          Ticketing System
        </Link>
        <div className="space-x-4">
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
