import { useState } from 'react';
import './index.scss';

export const FilterFieldDropdown: React.FC<{ props: { label: string; src: string } }> = ({ props }) => {
    const { label, src } = props;
    const [shouldDropdownShow, handleShowing] = useState(false);
    return (
        <div
            className="filter-item"
            onClick={() => handleShowing(prev => !prev)}
        >
            <span className="filter-item__title">
                {label}
            </span>
            <img
                className="filter-item__picture"
                src={src}
                alt={src && "filter icon"}
            />
            <div
                className="filter-item__dropdown"
                style={{ display: shouldDropdownShow ? 'block' : 'none' }}
            >

            </div>
        </div>
    )
}
