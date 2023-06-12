// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { useParams } from 'react-router';

import { FootballerCardIllustrationsRadar } from '@/app/components/common/Card/CardIllustrationsRadar';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { PlayerCard } from '@/app/components/common/PlayerCard';
import { ToastNotifications } from '@/notifications/service';
import { RootState } from '@/app/store';
import { openUserCard } from '@/app/store/actions/cards';
import { setCurrentUser } from '@/app/store/actions/users';
import WalletService from '@/wallet/service';
import { Sell } from '@/app/components/common/Card/popUps/Sell';
import { CasperNetworkClient } from '@/api/casper';
import { CasperNetworkService } from '@/casper/service';

import CardPageBackground from '@static/img/FootballerCardPage/background.png';
import backButton from '@static/img/FootballerCardPage/back-button.png';

import '../index.scss';

const Card: React.FC = () => {
    const dispatch = useDispatch();

    const [isMinted, setIsMinted] = useState<boolean>(false);
    const user = useSelector((state: RootState) => state.usersReducer.user);
    const { card } = useSelector((state: RootState) => state.cardsReducer);
    const { id }: { id: string } = useParams();

    const [isOpenSellPopup, setIsOpenSellPopup] = useState<boolean>(false);

    const casperClient = new CasperNetworkClient();
    const casperService = new CasperNetworkService(casperClient);

    /** Handle opening of a selles pop-up. */
    const handleOpenSellPopup = () => {
        setIsOpenSellPopup(true);
    };

    /** implements opening new card */
    async function openCard() {
        try {
            await dispatch(openUserCard(id));
        } catch (error: any) {
            ToastNotifications.couldNotOpenCard();
        }
    }

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    /** mints a card */
    const mint = async() => {
        try {
            const walletService = new WalletService(user);
            await walletService.mintNft(id);
        }
        catch (e: any) {
            ToastNotifications.somethingWentsWrong();
        }
    };

    /** approves nft minting */
    const approve = async() => {
        try {
            const approveData = await casperService.approve(card.id);

            const walletService = new WalletService(user);
            await walletService.approveNftMint(approveData);
        } catch (e: any) {
            ToastNotifications.somethingWentsWrong();
        }
    };

    useEffect(() => {
        setUser();
        openCard();
    }, []);

    return (
        card &&
        <>{isOpenSellPopup && <Sell setIsOpenSellPopup={setIsOpenSellPopup} />}
            <div className="card">
                <div className="card__wrapper">
                    <div className="card__back">
                        <Link className="card__back__button" to="/cards">
                            <img src={backButton} alt="back-button" className="card__back__button__image" />
                            Back
                        </Link>
                    </div>
                    <div className="card__info">
                        <PlayerCard className="card__player" id={card.id} />
                        <div className="card__player__info">
                            <h2 className="card__name">{card.playerName}</h2>
                            <div className="card__mint-info">
                                <div className="card__mint-info__nft">
                                    <span className="card__mint-info__nft-title">NFT:</span>
                                    <div className="card__mint-info__nft__content">
                                        <span className="card__mint-info__nft-value">
                                            {isMinted ? 'minted to Polygon' : 'not minted'}
                                        </span>
                                        {!isMinted &&
                                            <>
                                                <button className="card__mint" onClick={mint}>
                                                    Mint now
                                                </button>
                                                <button className="card__mint" onClick={approve}>
                                                    Approve
                                                </button>
                                            </>
                                        }
                                    </div>
                                </div>
                                <div className="card__mint-info__club">
                                    <span className="card__mint-info__club-title">Club:</span>
                                    <span className="card__mint-info__club-name">FC228</span>
                                </div>
                                <button
                                    className="card__sell-btn"
                                    onClick={handleOpenSellPopup}
                                >
                                    <span className="card__sell-btn__text">SELL</span>
                                </button>
                            </div>
                        </div>
                        <div className="card__illustrator-radar">
                            <h2 className="card__illustrator-radar__title">Skills</h2>
                            <FootballerCardIllustrationsRadar card={card} />
                        </div>
                    </div>
                    <div className="card__stats-area">
                        <FootballerCardStatsArea card={card} />
                    </div>
                </div>
                <img src={CardPageBackground} alt="background" className="card__bg" />
            </div>
        </>
    );
};

export default Card;
