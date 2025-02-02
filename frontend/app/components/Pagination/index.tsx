import React from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';

interface PaginationProps {
	currentPage: number;
	totalPages: number;
	onPageChange: (page: number) => void;
	isLoading: boolean;
}

const Pagination = ({ currentPage, totalPages, onPageChange, isLoading }: PaginationProps) => {
	const pages = Array.from({ length: totalPages }, (_, i) => i + 1);

	if (totalPages <= 1) return null;

	return (
		<div className="flex items-center justify-center space-x-2 mt-6">
			<button
				onClick={() => onPageChange(currentPage - 1)}
				disabled={currentPage === 1 || isLoading}
				className="p-2 rounded-md hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
			>
				<ChevronLeft className="w-5 h-5" />
			</button>

			{pages.map((page) => (
				<button
					key={page}
					onClick={() => onPageChange(page)}
					disabled={isLoading}
					className={`px-3 py-1 rounded-md ${
						currentPage === page
							? 'bg-blue-500 text-white'
							: 'hover:bg-gray-100'
					} disabled:opacity-50 disabled:cursor-not-allowed`}
				>
					{page}
				</button>
			))}

			<button
				onClick={() => onPageChange(currentPage + 1)}
				disabled={currentPage === totalPages || isLoading}
				className="p-2 rounded-md hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
			>
				<ChevronRight className="w-5 h-5" />
			</button>
		</div>
	);
};

export default Pagination;
