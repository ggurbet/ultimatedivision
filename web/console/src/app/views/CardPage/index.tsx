// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { toast } from 'react-toastify';
import { useParams } from 'react-router';
import MetaMaskOnboarding from '@metamask/onboarding';

import { FootballerCardIllustrationsRadar } from '@/app/components/common/Card/CardIllustrationsRadar';
import { FootballerCardPrice } from '@/app/components/common/Card/CardPrice';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { PlayerCard } from '@/app/components/common/PlayerCard';

import { RootState } from '@/app/store';
import { openUserCard } from '@/app/store/actions/cards';
import { ServicePlugin } from '@/app/plugins/service';
import { setCurrentUser } from '@/app/store/actions/users';
import { metamaskNotifications } from '@/app/internal/notifications';

import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';
import CasperTransactionService from '@/casper';

import CardPageBackground from '@static/img/FootballerCardPage/background.png';
import backButton from '@static/img/FootballerCardPage/back-button.png';

import './index.scss';

const VELAS_WALLET_TYPE = 'velas_wallet_address';
const CASPER_WALLET_TYPE = 'casper_wallet_address';
const METAMASK_WALLET_TYPE = 'wallet_address';

const Card: React.FC = () => {
    const dispatch = useDispatch();

    const [isMinted, setIsMinted] = useState<boolean>(false);

    const user = useSelector((state: RootState) => state.usersReducer.user);
    const { card } = useSelector((state: RootState) => state.cardsReducer);
    const { id }: { id: string } = useParams();

    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const service = ServicePlugin.create();

    const usersClient = new UsersClient();
    const usersService = new UsersService(usersClient);

    /** implements opening new card */
    async function openCard() {
        try {
            await dispatch(openUserCard(id));
        } catch (error: any) {
            /** TODO: it will be reworked with notification system */
        }
    }

    /** implements opening new card */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }

    /** Mints chosed card with velas */
    const velasMint = async() => { };

    /** Mints chosed card with metamask */
    const metamaskMint = async() => {
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore .
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });
                await service.sendTransaction(id);
            } catch (error: any) {
                metamaskNotifications(error);
            }
        } else {
            onboarding.startOnboarding();
        }
    };

    /** Mints chosed card with casper */
    const casperMint = async() => {
        try {
            const casperTransactionService = new CasperTransactionService(user.casperWallet);

            await casperTransactionService.mint(id);
        } catch (error: any) {
            toast.error(`${error}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    const mint = async() => {
        try {
            const user = await usersService.getUser();

            const mintingNfts = new Map();

            const walletMintingTypes = [
                {
                    walletType: VELAS_WALLET_TYPE,
                    mint: velasMint,
                },
                {
                    walletType: CASPER_WALLET_TYPE,
                    mint: casperMint,
                },
                {
                    walletType: METAMASK_WALLET_TYPE,
                    mint: metamaskMint,
                },
            ];

            walletMintingTypes.forEach(walletMintingType =>
                mintingNfts.set(walletMintingType.walletType, walletMintingType.mint));

            mintingNfts.get(user.walletType)();
        } catch (e) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    useEffect(() => {
        setUser();
        openCard();
    }, []);

    return (
        card &&
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
                                        <button className="card__mint" onClick={mint}>
                                            Mint now
                                        </button>
                                    }
                                </div>
                            </div>
                            <div className="card__mint-info__club">
                                <span className="card__mint-info__club-title">Club:</span>
                                <span className="card__mint-info__club-name">FC228</span>
                            </div>
                            <FootballerCardPrice card={card} isMinted={isMinted} />
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
    );
};

export default Card;
