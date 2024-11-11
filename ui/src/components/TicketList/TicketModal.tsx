import React from 'react';
import { Ticket } from '../../types/ticketTypes';

interface TicketModalProps {
  ticket: Ticket;
  onClose: () => void;
}

const TicketModal: React.FC<TicketModalProps> = ({ ticket, onClose }) => {
  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
      <div className="bg-white p-6 rounded-lg max-w-lg w-full">
        <h3 className="text-2xl font-semibold mb-4">Ticket Details</h3>
        <p>
          <strong>Name:</strong> {ticket.name}
        </p>
        <p>
          <strong>Description:</strong> {ticket.desc}
        </p>
        <p>
          <strong>Allocation:</strong> {ticket.allocation}
        </p>
        <div className="mt-4 flex justify-end">
          <button
            type="button"
            onClick={onClose}
            className="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
};

export default TicketModal;
