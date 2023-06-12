// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link, useParams } from 'react-router-dom';

import { RootState } from '@/app/store';
import { openMarketplaceCard } from '@/app/store/actions/marketplace';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { BidArea } from '@/app/components/common/Card/BidArea';
import { FootballerCardIllustrationsRadar } from '@/app/components/common/Card/CardIllustrationsRadar';
import { PlayerCard } from '@/app/components/common/PlayerCard';
import { ToastNotifications } from '@/notifications/service';
import { setCurrentUser } from '@/app/store/actions/users';

import CardPageBackground from '@static/img/FootballerCardPage/background.png';
import backButton from '@static/img/FootballerCardPage/back-button.png';

import '../index.scss';

const Lot: React.FC = () => {
    const dispatch = useDispatch();
    const { lot } = useSelector((state: RootState) => state.marketplaceReducer);

    const { id }: { id: string } = useParams();

    /** implements opening new card */
    async function openLot() {
        try {
            await dispatch(openMarketplaceCard(id));
        } catch (error: any) {
            ToastNotifications.couldNotOpenCard();
        };
    };

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    useEffect(() => {
        openLot();
        setUser();
    }, []);

    return (
        lot &&
        <div className="lot">
            <div className="lot__wrapper">
                <div className="lot__back">
                    <Link className="lot__back__button" to="/marketplace">
                        <img src={backButton} alt="back-button" className="lot__back__button__image" />
                        Back
                    </Link>
                </div>
                <div className="lot__info">
                    <PlayerCard className="lot__player" id={lot.card.id} />
                    <div className="lot__player__info">
                        <h2 className="lot__name">{lot.card.playerName}</h2>
                        <div className="lot__mint-info">
                            <div className="lot__mint-info__nft">
                                <span className="lot__mint-info__nft-title">NFT:</span>
                                <div className="lot__mint-info__nft__content">
                                    <span className="lot__mint-info__nft-value">
                                        'minted to Polygon'
                                    </span>
                                </div>
                            </div>
                            <div className="lot__mint-info__club">
                                <span className="lot__mint-info__club-title">Club:</span>
                                <span className="lot__mint-info__club-name">FC228</span>
                            </div>
                            <BidArea />
                        </div>
                    </div>
                    <div className="lot__illustrator-radar">
                        <h2 className="lot__illustrator-radar__title">Skills</h2>
                        <FootballerCardIllustrationsRadar card={lot.card} />
                    </div>
                </div>
                <div className="lot__stats-area">
                    <FootballerCardStatsArea card={lot.card} />
                </div>
            </div>
            <img src={CardPageBackground} alt="background" className="lot__bg" />
        </div>

    );
};

export default Lot;
