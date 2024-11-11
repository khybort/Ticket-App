import React from 'react';
import { Ticket } from '../../types/ticketTypes';

interface TicketItemProps {
  ticket: Ticket;
  onGetTicket: (id: number) => void;
  quantities: { [key: number]: number };
  onQuantityChange: (id: number, value: number) => void;
  onPurchase: (id: number, quantity: number) => void;
  loading: boolean;
}

const TicketItem: React.FC<TicketItemProps> = ({
  ticket,
  onGetTicket,
  quantities,
  onQuantityChange,
  onPurchase,
  loading,
}) => {
  return (
    <li
      key={ticket.id}
      className={`flex justify-between items-start p-4 border rounded-lg ${
        ticket.allocation > 0 ? 'bg-green-100' : 'bg-red-100'
      }`}
    >
      <div className="flex flex-col">
        <span
          className="text-lg font-medium cursor-pointer"
          onClick={() => onGetTicket(ticket.id)}
          tabIndex={0}
          role="button"
          onKeyDown={(e) => e.key === 'Enter' && onGetTicket(ticket.id)}
        >
          {ticket.name}
        </span>
        <p className="text-sm text-gray-500">{ticket.desc}</p>
      </div>
      <div className="flex flex-col items-end">
        <span className="text-xl font-semibold">
          {ticket.allocation} available
        </span>
        <span
          className={`px-4 py-1 rounded-full text-white font-semibold ${
            ticket.allocation > 0 ? 'bg-green-500' : 'bg-red-500'
          }`}
        >
          {ticket.allocation > 0 ? 'Available' : 'Sold Out'}
        </span>

        {ticket.allocation > 0 && (
          <div className="flex flex-col items-end mt-4">
            <input
              type="number"
              min="1"
              max={ticket.allocation}
              value={quantities[ticket.id] || 1}
              onChange={(e) =>
                onQuantityChange(ticket.id, Number(e.target.value))
              }
              className="px-3 py-2 border rounded-lg mb-2 w-24 text-center"
            />
            <button
              type="button"
              onClick={() => onPurchase(ticket.id, quantities[ticket.id] || 1)}
              disabled={
                loading ||
                quantities[ticket.id] <= 0 ||
                quantities[ticket.id] > ticket.allocation
              }
              className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 focus:outline-none disabled:opacity-50"
            >
              {loading ? 'Processing...' : 'Purchase'}
            </button>
          </div>
        )}
      </div>
    </li>
  );
};

export default TicketItem;
