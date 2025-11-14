package workflow

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"web/gopkg/log"
)

type ChatStream struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    DataItem `json:"data"`
}

type DataItem struct {
	Results any `json:"results,omitempty"`
}

type MessageItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func duration() time.Duration {
	return time.Second * 720
}

func (s *Service) ChatStream(ctx context.Context, question string) (chan ChatStream, error) {
	var logger = log.SugarContext(ctx)
	var response = make(chan ChatStream, 1000)
	go func() {
		ctx, cancel := context.WithTimeout(context.WithValue(context.Background(), "x-request-id", ctx.Value("x-request-id")), duration())
		defer func() {
			if err := recover(); err != nil {
				log.SugarContext(ctx).Errorw("流式问答问答异常", "error", err)
			}
			close(response)
			cancel()
		}()

		//page, count, err := ScanByPage(documentID, nil, []int{}, 1, 500)
		//if err != nil || count == 0 {
		//	log.SugarContext(ctx).Errorw("Markdown解析失败", "error", err)
		//	return
		//}
		//
		//index := 0
		//fullTextList := make([]string, 0)
		//for _, v := range page {
		//	content := v.Content
		//	if content == "" {
		//		content = v.Title
		//	}
		//	index++
		//	fullTextList = append(fullTextList, content)
		//}
		//
		//documentInfo := strings.Join(fullTextList, " \n")
		//documentInfoSubText := documentInfo[:20000]
		documentInfoSubText := `POLITECNICO DI TORINO Repository ISTITUZIONALE \nModeling the lithium loop in a liquid metal pool-type divertor \nOriginal Modeling the lithium loop in a liquid metal pool-type divertor / Nallo, GIUSEPPE FRANCESCO; Carli, Stefano; Caruso, G.; Crisanti, F.; Mazzitelli, G.; Savoldi, Laura; Subba, Fabio; Zanino, Roberto. - In: FUSION ENGINEERING AND DESIGN. - ISSN 0920-3796. - ELETTRONICO. - 125:(2017), pp. 206-215. [10.1016/j.fusengdes.2017.07.004] \nAvailability: This version is available at: 11583/2679206 since: 2021-07-02T12:06:14Z \nPublisher: Elsevier Ltd \nPublished DOI:10.1016/j.fusengdes.2017.07.004 \nTerms of use: \nThis article is made available under terms and conditions as specified in the  corresponding bibliographic description in the repository \n(Article begins on next page) \nModeling the Lithium Loop in the DTT Liquid Metal Divertor \nG. F. Nalloa, S. Carlia, G. Carusob, F. Crisantic, G. Mazzitellic, L. Savoldia, F. Subbaa, R. Zaninoa aNEMO group, Dipartimento Energia, Politecnico di Torino, Torino, Italy $b$ Università degli Studi di Roma “La Sapienza”, Roma, Italy cENEA Frascati, Italy \nSolutions for the steady-state power exhaust problem in future fusion reactors (e.g. DEMO) are neither provided by present experiments nor will be by ITER, because the expected heat fluxes, as well as the level of neutron irradiation, will be much higher. Dedicated work packages are being devoted to this problem within EUROfusion and even a dedicated facility (the Divertor Tokamak Test - DTT) is being proposed in Italy. Among the possible solutions to the problem, a liquid metal (LM) divertor was proposed more than 20 years ago. The particularly attractive feature of this solution is the absence of damage to the wall, even in the case of large heat fluxes, thanks to the high latent heat of evaporation and to the liquid nature of the wall, which can be constantly replenished. The present work aims at developing a simple model of the LM loop including the most important physical phenomena and allowing to roughly determine the operating range of the system, in terms of surface temperatures and vapor pressures. This work therefore sets the basis for the conceptual design of a LM divertor for the DTT facility. The preliminary model has been set up including the incoming plasma heat load and a basic treatment of the interactions of Li vapor with the plasma. The reduction of the Li vapor efflux due to ionization by the plasma is also taken into account. The model includes two chambers: a first divertor box, the evaporation chamber (EC), is open towards a second divertor box, the differential chamber (DC), which is in turn connected to the main plasma chamber (MC). \nThe model is used to study the effectiveness of the LM vapor in radiating isotropically the parallel heat flux incoming in the divertor. The results indicate that the presence of the DC allows a significant reduction of the Li vapor efflux towards the MC, which in turn would imply a lower contamination of the core plasma. Future studies will include a capillary-porous structure (CPS) coating of the internal surfaces of the divertor and a 2D approach to both plasma and Li vapor modeling. \nKeywords: Plasma Exhaust, Divertor, Lithium, SOL, DTT \n1. Introduction \nSafe power exhaust, even in steady state, is one of the major issues in fusion reactors and a potential showstopper towards the production of the first kWh from the fusion energy source, which the EU roadmap [1] has set as the major target for its DEMO reactor in 2050. The power produced by deuterium-tritium reactions in the alpha particles channel may be partly -- more or less isotropically -- radiated, but the rest reaches the plasmafacing components (PFCs) in the strongly anisotropic channel of plasma advection-conduction. This leads to high particle and heat fluxes, because of the relatively small wetted areas associated with the thin scrape-off layer (SOL) predicted in future machines, affecting not only the PFCs lifetime but also the core plasma purity (measured by the $Z_{\\mathrm{eff}}$ parameter), because of sputtering. \nIn ITER, the control of the steady-state peak heat load qpeak below $10\\mathrm{MW/m}^{2}$ on the PFCs, as well as of the $Z_{\\mathrm{eff}}.$ relies on a single-null divertor with W target and on the Be first wall (FW), combined with detached plasma operation and seed impurity puffing [2, 3]. Even if this complex combination of conditions should be confirmed experimentally in ITER, extrapolation to DEMO is not automatically guaranteed. Furthermore, as the increase in size from ITER to DEMO is much smaller than the increase in the design thermal power to be produced by the two machines, in DEMO it will be even more difficult to meet the technological constraints on $\\mathbf{q}_{\\mathrm{peak}}$ related to the use of a solid divertor. \nAmong the risk-mitigation strategies currently foreseen, the one especially relevant for the present work is the liquid metal (LM) divertor [4]: the LM evaporation, together with the plasma cooling by interactions with the LM vapor, could in principle guarantee the exhaust of hundreds of $\\mathrm{{MW/m}}^{2}$ , with much more limited, if any, damage to the target than the solid target option, and consequent increase of the divertor lifetime. \nPossible LM choices include in the first place Li [2], which will be considered here, but also others, e.g. Sn. As to the nature of the LM target, different options are being considered, ranging in complexity from a simple pool to a moving liquid film, to the use of a so-called capillary porous structure (CPS) [4], recently tested on the liquid Li limiter (LLL) in FTU [5]. While the CPS should guarantee, better than other solutions, the avoidance of splashing phenomena with generation of LM droplets, which could easily compromise the plasma purity, we will refer in the present work to the simpler case of a LM pool without CPS and without external (pumped) circulation of the LM. \nIn the EU, significant attention is being given to these problems within the EUROfusion Work Packages DTT1 and DTT2. DTT is also the name of an Italian proposal for a machine entirely devoted to the issues of power exhaust and $Z_{\\mathrm{eff}}$ in DEMO perspective [6]. \n2. System description \nStarting from the current design of the DTT chamber [7], a first preliminary sketch of a possible liquid metal divertor geometry to fit in the available space has been prepared, see figure 1. This is based on the idea originally proposed by Nagayama [8] and eventually further developed in [9]: the SOL plasma flowing from the main plasma chamber in a reference Single Null (SN) DTT equilibrium enters first the Differential Chamber (DC, see figure 1) and finally the Evaporation Chamber (EC), where the LM pool is located. \n!(Image)(s3://studio-1255000019/dataset-service-develop/f2dfd92d-21e6-40aa-8da4-2c76a3bcd4e5/047cfbf5-f030-4789-abfa-7bd22acdc954/Modeling the lithium loop in a liquid metal pool-type divertor_attachment/603ca4a2dde1782e425c19385efdcb4a030ece6f9b2593b7911e95bb4788b9c2.jpg)Fig. 1: (a) DTT main plasma chamber with the divertor highlighted and (b) preliminary sketch of the LM divertor. \nThe schematic representation of the EC used to write down the model equations is shown in figure 2. Even though the shape shown in the schematic is different with respect to the one in figure 1, this is not relevant for a 0D model, the only requirement being to conserve the surface areas and the chamber volumes. The sub-systems are: \nthe liquid Li in the pool;the Li vapor in the remaining volume of the EC; the Li vapor in the DC;the solid walls in contact with liquid Li (identified in the following with the subscript pool);the solid walls in contact with the Li vapor in the EC (identified in the following with subscript $E C)$ ).the solid walls in contact with the Li vapor in the DC (identified in the following with subscript $D C)$ . \n3. Phenomenology \nThe Li evaporating from the pool flows upwards, then either condenses on the relatively colder surfaces of the EC or moves to the DC, where it either condenses or moves outside of the box-structure of the divertor towards the main plasma chamber. The condensed Li is assumed to flow back to the Li pool both from the EC (by gravity) and from the DC (by means of an external circuit, not included in this model for the time being). \n!(Image)(s3://studio-1255000019/dataset-service-develop/f2dfd92d-21e6-40aa-8da4-2c76a3bcd4e5/047cfbf5-f030-4789-abfa-7bd22acdc954/Modeling the lithium loop in a liquid metal pool-type divertor_attachment/9c58cf175ea2659f2174cfe91dfac7682488483ffabcd2cb5998a01ae561bc2e.jpg)Fig. 2: Schematic of the computational domain. \nThe presence of a DC in the original Nagayama proposal is motivated by the necessity to reduce the core plasma contamination associated with the eroded (evaporated/sputtered) Li flowing out of the EC. The extra chamber allows for differential pumping, i.e. the connection of two chambers having different pressures by means of intermediate chambers, actively and/or passively pumped (see figure 3). In such a system, when a large pressure difference is involved, choked flow is likely to occur between successive chambers [10]. \n!(Image)(s3://studio-1255000019/dataset-service-develop/f2dfd92d-21e6-40aa-8da4-2c76a3bcd4e5/047cfbf5-f030-4789-abfa-7bd22acdc954/Modeling the lithium loop in a liquid metal pool-type divertor_attachment/470771a5cb8f800512cbae8968f5d7ac416f185cf8e9e21016cf25f83118d5af.jpg)Fig. 3: schematic of a differentially pumped system. \nThe intermediate chambers allow a progressive reduction in mass flow rate from the higher pressure boxes to the lower pressure boxes, thanks to the – active and/or passive – pumping of the vapor. In the concept considered here, this is achieved by means of net condensation of Li vapor on the walls of the EC and of the DC, i.e. by a passive pumping mechanism. An active differential pumping based on turbomolecular pumps could also be foreseen, in order to differentially pump non-condensable gases. \nAs it will be pointed out later, the presence of the SOL plasma further reduces the Li efflux between successive chambers (and, eventually, towards the main plasma chamber), thanks to ionization of Li vapor. \nEroded Li is readily ionized by the plasma due to its low first ionization energy $(\\sim5\\ \\mathrm{eV})$ , resulting in a significant reduction of the heat flux to the divertor strike point thanks to the related plasma cooling effect [11, 12]. This effect, which in the following will be referred to as the Li “vapor shield”, is a consequence of ionization, line radiation and continuum radiation due to interactions of the free electrons of the SOL plasma with the vapor and could lead to a more manageable flux of Li vapor towards the core plasma (a reduced heat load on the Li pool would reduce the evaporation rate). From the point of view of the Li mass balance, the fact that Li atoms get “entrained” by the SOL plasma implies a further reduction of the efflux from the EC and consequently from the DC to the main plasma chamber. \nThe chemical reactivity of liquid Li with hydrogenic species allows in principle the LM to capture fuel particles (increased pumping), allowing a low recycling operation in tokamaks [13], until the Li layer is fully saturated. This occurs because the thermally enhanced diffusion of D atoms in the LM allows in principle to avoid D atoms “piling up” at the surface. Detailed calculations on this point are beyond the scope of this work, but the work in [14] could be employed for a study of D diffusion under various plasma conditions. \nAn external liquid Li circulation loop for addressing tritium inventory and dust accumulation issues is essential for the steady state behavior of the system hereby analyzed. A detailed analysis of this external loop is beyond the scope of the present work, but recently Ono et al. ([15]) suggested that it is possible to effectively remove tritium and impurities from the liquid Li. \n4. Model description \nA simplified but self-consistent thermodynamic (0D) model is presented to compute the thermodynamic state of the Li liquid-vapor system, the average wall temperatures and the Li vapor efflux towards the main plasma chamber in an axisymmetric LM divertor of box type. Notwithstanding its simplicity, the model can be considered to be an important step, since a comprehensive analysis of such a system (i.e. a divertor based on the closed-box concept) cannot be found in the literature, to the best of our knowledge, the only pioneering work being the recently published model in [9, 16].  Compared to the latter, the present study employs a transient model in order to reach the steady state, does not assume choked flow between successive boxes and evaluates the average wall temperatures based on the walls energy balance rather than imposing it a priori. \nThe main assumptions in our model are as follows: \n1. The Li vapor is \na. approximated as an ideal gas. This is verified a posteriori to be acceptable, thanks to the very low pressure foreseen in the EC – see section 5;b. monoatomic (the fraction of $\\operatorname{Li}_{2}$ in vapor phase is verified a posteriori to be between 5 and $10\\%$ at the temperatures foreseen for this system [17]);c. optically thin with respect to radiation, i.e. the radiated power resulting from the interactions between the SOL plasma and the Li vapor is assumed to be deposited on the walls and on the pool surface without absorption within the \nmedium. This is justified due to the extremely low vapor densities expected in the system; \nd. collisional, i.e. the mean free path of the Li atoms in vapor phase is small compared to some characteristic length of the system. Further details concerning this assumption are given in Appendix 1;e. partly lost –as a neutral- towards the main plasma chamber and compensated by an equal amount of replenishing liquid Li supplied to the pool, i.e. the total mass of Li in the system is conserved. This simulates the effect of a Li reservoir that is often foreseen in similar systems [18];f. partly entrained (ionized) by the plasma and recombining before reaching the pool;g. flowing isenthalpically between neighboring boxes;h. in thermodynamic equilibrium with the liquid phase. \nThe Li pool \na. is optically thick with respect to radiation, i.e., it absorbs all radiated power directed towards it, whereas radiation emitted from it to the (colder) walls is negligible with respect to radiation from the Li vapor shield (checked a posteriori);b. receives a fraction $f\\approx A_{p o o l}/(A_{w,E C}+A_{p o o l})$ of the radiated power in the EC, whereas the remaining fraction $(1-f)$ is directed towards the walls;c. instantaneously collects the Li re-condensed on both the EC and DC walls, i.e., the dynamics of the condensed Li film and the external Li recirculation circuit are not included in the model for the time being. \nThe effect of the particle influx from the plasma is neglected, i.e. no account is given here about retention (the properties of the Li pool are not modified), sputtering (no sputtering source has been included in the mass conservation equations for the Li vapor) and pressure build-up due to noncondensable gases (D, T, He). \nNo radiated power is deposited on the DC walls. \nThe conservation of mass for the Li in the EC (detailed scheme shown in figure 4 and control volume shown in figure 5 (a)) is described by \n$$\\begin{array}{r l}&{\\frac{d N_{L i,E C}}{d t}=\\dot{N}{r e c o l l e c t i o n,E C}+\\dot{N}{r e p l}+\\dot{N}{r e c o m b,E C}-}\\\\ &{\\dot{N}{n o z,E C\\rightarrow D C,n o e n t r}-\\dot{N}{e n t r,E C}-\\dot{N}{r e c o n d,E C}}\\end{array}$$ \nwhere: \n$N_{L i,E C}$ is the total number of Li atoms in the EC; $\\dot{N}{r e c o l l e c t i o n,E C}=\\dot{N}{r e c o n d,E C}+\\dot{N}{r e c o n d,D C}$ is the particle flow rate of recondensed Li atoms on the walls of both chambers, which are assumed to be instantaneously returned to the pool according to assumption 2c. In particular, $\\dot{N}{r e c o n d,E C}$ (atoms/s) is the net re-condensation rate on the walls of the EC, which is calculated by means of a modified HertzKnudsen equation [19] \n$$\\begin{array}{r l}&{\\dot{N}{r e c o n d,E C}=\\dot{N}{c o n d,E C}-\\dot{N}{e v,E C}=\\eta\\cdot A{w,E C}\\cdot10^{3}\\cdot}\\\\ &{\\left(\\frac{p_{E C}\\cdot N_{A v}}{\\sqrt{2\\pi m_{L i}R_{0}T_{E C}}}-\\frac{p_{s a t}(T_{E C})\\cdot N_{A v}}{\\sqrt{2\\pi m_{L i}R_{0}T_{w,E C}}}\\right)}\\end{array}$$ \nwhere $\\dot{N}{c o n d,E C}$ (atoms/s) is the condensation rate to the walls of the EC, $\\dot{N}{e v,E C}$ (atoms/s) is the evaporation rate from the EC, T $T_{w,E C}$ (K) is the temperature of the EC walls, $T_{E C}$ (K) is the temperature of the Li vapor in the EC, $p_{s a t}(T_{E C})$ (Pa) is the saturation pressure evaluated at $T_{E C},p_{E C}$ (Pa) is the pressure of Li vapor. $\\eta$ is an empirical coefficient estimated in [19] to be equal to 1.66, $A_{w,E C}(\\mathrm{m}^{2})$ is the surface area of the EC walls facing the Li vapor, $m_{L i}$ $\\mathrm{(g/mol)}$ is the molar mass of Li, $N_{A v}$ is the Avogadro number and $R_{0}$ $\\mathrm{(J/kmol/K)}$ is the universal gas constant. (A more detailed analysis including an energy balance for the Li film would be more appropriate, but this is left for future work). The same approach is employed for evaluating 𝑁̇𝑟𝑒𝑐𝑜𝑛𝑑,𝐷𝐶.. \n$\\dot{N}{r e c o m b,E C}=\\dot{N}{e n t r,E C}+\\dot{N}{e n t r,E C\\rightarrow D C}+\\dot{N}{e n t r,D C}+$ $\\dot{N}{e n t r,D C\\rightarrow M C}$ is the particle flow rate of Li atoms entrained by the plasma in the entire system, which are all assumed to recombine within the EC. Indeed, as recombination is assumed to occur in the EC, from the Li particle balance point of view, it means that also atoms ionized in the DC enter the particle balance within the EC as a source term. In particular, $\\dot{N}{e n t r,E C}$ (atoms/s) is the particle flow rate of Li vapor entrained by the plasma within the EC, evaluated relying on the statistical mechanics formulation of the particle flux striking on a surface (Langmuir flux), which assumes a Maxwellian distribution of the atoms. The nature of this expression is therefore exactly the same as the original Hertz-Knudsen one (i.e. equation (2) without the $\\eta$ factor), but for a “purely condensing wall” [10], since the plasma cannot release entrained atoms until recombination has occurred: locally, it acts as a perfect particle sink. The area employed in this formulation is the outer surface of the SOL plasma in the EC. $\\dot{N}_{e n t r,D C}$ is evaluated in the same way. \n$\\dot{N}_{e n t r,E C\\rightarrow D C}$ (atoms/s) is the particle flow rate of Li vapor entrained by the plasma while passing from EC to DC, evaluated as: \n$$\\begin{array}{r}{\\Dot{N}{e n t r,E C\\rightarrow D C}=\\Dot{N}{n o z,E C\\rightarrow D C,n o e n t r}\\cdot\\left(\\frac{\\lambda_{p,O M P}}{A_{n o z}}\\cdot f_{E X P}\\right)}\\end{array}$$ \nwhere $\\dot{N}_{n o z,E C\\rightarrow D C,n o e n t r}$ (atoms/s) is the Li vapor flow rate from the EC to the DC through the aperture (nozzle) between the two, evaluated assuming isenthalpic flow and neglecting the presence of the plasma: \n$$\\begin{array}{r}{\\dot{N}{n o z,E C\\rightarrow D C,n o e n t r}=\\bigg[\\rho{n o z}\\cdot A_{n o z}\\cdot}\\\\ {\\sqrt{2\\cdot\\frac{\\gamma}{\\gamma-1}\\cdot\\left(\\frac{p_{E C}}{\\rho_{E C}}-\\frac{p_{n o z}}{\\rho_{n o z}}\\right)}\\bigg]\\cdot\\frac{N_{A v}}{m_{L i}}\\cdot10^{3}}\\end{array}$$ \nwhere $\\rho~\\mathrm{(kg/m^{3})}$ is a density, $A_{n o z}\\mathrm{~\\ensuremath~{~(m^{2})}~}$ is the passage area between boxes, $\\gamma=5/3$ is the isentropic exponent and the subscript noz refers to gas conditions at the nozzle. In particular: \nif 𝑝𝐷𝐶 $\\begin{array}{r}{\\frac{p_{D C}}{p_{E C}}\\leq\\Big(\\frac{2}{\\gamma+1}\\Big)^{\\frac{\\gamma}{\\gamma-1}}}\\end{array}$ , then $\\begin{array}{r}{p_{n o z}=p_{E C}\\left(\\frac{2}{\\gamma+1}\\right)^{\\frac{`

		MessageList := make([]MessageItem, 0)
		MessageList = append(MessageList, MessageItem{
			Role:    "user",
			Content: question,
		})
		payload := NewPayload().
			SetWorkflowID("e367c063-f9a2-5479-85a6-63928f48d84b").
			SetChatMessages(MessageList).
			SetChatFullText(documentInfoSubText)

		// 以SSE方式请求并返回流
		apiUrl := "http://10.122.9.26:31373/v1/openapi/workflow/running/event_stream"

		body, err := json.Marshal(payload)
		if err != nil {
			logger.Errorw("marshal payload error", "error", err)
			return
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiUrl, bytes.NewReader(body))
		if err != nil {
			logger.Errorw("new request error", "error", err)
			return
		}

		client := &http.Client{}
		httpResp, err := client.Do(req)
		if err != nil {
			logger.Errorw("do request error", "error", err)
			return
		}
		defer httpResp.Body.Close()

		reader := bufio.NewReader(httpResp.Body)
		var (
			eventName   string
			dataBuilder strings.Builder
		)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() != "EOF" {
					logger.Errorw("read sse line error", "error", err)
				}
				break
			}
			line = strings.TrimRight(line, "\r\n")
			if line == "" { // 事件分隔
				if eventName != "" && dataBuilder.Len() > 0 {
					var sseResp Response
					if err := json.Unmarshal([]byte(strings.TrimSpace(dataBuilder.String())), &sseResp); err == nil {
						var results any
						if m, ok := sseResp.Data.(map[string]any); ok {
							if r, ok := m["results"]; ok {
								results = r
							} else {
								results = m
							}
						} else {
							results = sseResp.Data
						}
						response <- ChatStream{
							Code:    sseResp.Code,
							Message: sseResp.Msg,
							Data:    DataItem{Results: results},
						}
					}
				}
				eventName = ""
				dataBuilder.Reset()
				continue
			}
			if strings.HasPrefix(line, "event:") {
				eventName = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
				continue
			}
			if strings.HasPrefix(line, "data:") {
				dataBuilder.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "data:")))
				dataBuilder.WriteByte('\n')
				continue
			}
		}

		if eventName != "" && dataBuilder.Len() > 0 {
			var sseResp Response
			if err := json.Unmarshal([]byte(strings.TrimSpace(dataBuilder.String())), &sseResp); err == nil {
				var results any
				if m, ok := sseResp.Data.(map[string]any); ok {
					if r, ok := m["results"]; ok {
						results = r
					} else {
						results = m
					}
				} else {
					results = sseResp.Data
				}
				response <- ChatStream{
					Code:    sseResp.Code,
					Message: sseResp.Msg,
					Data:    DataItem{Results: results},
				}
			}
		}
	}()
	return response, nil
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
