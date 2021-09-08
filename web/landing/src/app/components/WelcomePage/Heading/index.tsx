// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { HeadingBunner }
    from '@components/WelcomePage/Heading/HeadingBunner';
import { HeadingInformation }
    from '@components/WelcomePage/Heading/HeadingInformation';
import { HeadingButton }
    from '@components/WelcomePage/Heading/HeadingButton';
import { HeadingModal }
    from '@components/WelcomePage/Heading/HeadingModal';

import { RegistrationPopup } from '@/app/views/RegistrationPopup';
import './index.scss';

export const Heading: React.FC = () => {
    const [offsetTop, handlescroll] = useState(0);

    const [showModal, setShowModal] = useState(false);
    /** just for test rigth now */
    const [showPopUp, setShowPopUp] = useState(false);
    const handlePopUp = () => setShowPopUp(!showPopUp);
    const closeModalOnPlay = () => {
        setShowModal(false);
        handlePopUp();
    };
    const showModalOnPlay = () => setShowModal(true);

    /** TODO: fix no-restricted-globals */
    useEffect(() => {
        document.addEventListener('scroll', () => {
            const bodyTop = document.body.getBoundingClientRect().top;

            handlescroll(bodyTop);
        });
    });

    return (
        <section className="heading" onScroll={() => scroll}>
            <HeadingBunner offsetTop={offsetTop} />
            <HeadingInformation />
            <HeadingButton
                handleShowModal={showModalOnPlay} />
            {showModal
                && <HeadingModal
                    handleCloseModal={closeModalOnPlay} />}
            {showPopUp && <RegistrationPopup handlePopUp={handlePopUp} />}
        </section>
    );
};
