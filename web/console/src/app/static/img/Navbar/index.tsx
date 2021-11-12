/** Dropdown Icon for mobile */
export function DropdownIcon() {
    return (
        <svg
            width="28"
            height="28"
            viewBox="0 0 28 28"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <rect width="28" height="6" fill="white" />
            <rect y="11" width="28" height="6" fill="white" />
            <rect y="22" width="28" height="6" fill="white" />
        </svg>
    );
};

/** Dropdown close icon */
export function CloseDropdownIcon() {
    return (
        <svg
            className="cross-menu-svg"
            width="40" height="40"
            viewBox="0 0 40 40"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <rect
                x="32.7266"
                y="3.63672"
                width="3.63636"
                height="40.9091"
                transform="rotate(45 32.7266 3.63672)"
                fill="white"
            />
            <rect
                x="35.4541"
                y="32.7275"
                width="3.63636"
                height="40.9091"
                transform="rotate(135 35.4541 32.7275)"
                fill="white"
            />
        </svg>
    );
};
