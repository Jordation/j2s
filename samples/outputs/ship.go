package main

type Summary struct {
	Data *Data `json:"data"`
	Meta *Meta `json:"meta"`
}

type Acceleration struct {
	ManeuveringG float64 `json:"maneuvering_g"`
	Main float64 `json:"main"`
	Retro float64 `json:"retro"`
	Vtol int `json:"vtol"`
	Maneuvering float64 `json:"maneuvering"`
	MainG float64 `json:"main_g"`
	RetroG float64 `json:"retro_g"`
	VtolG int `json:"vtol_g"`
}

type Agility struct {
	Roll int `json:"roll"`
	Acceleration *Acceleration `json:"acceleration"`
	Pitch int `json:"pitch"`
	Yaw int `json:"yaw"`
}

type Ammunition struct {
	Lifetime float64 `json:"lifetime"`
	Speed int `json:"speed"`
	Range int `json:"range"`
	Piercability *Piercability `json:"piercability"`
	DamageFalloffs *DamageFalloffs `json:"damage_falloffs"`
	Uuid string `json:"uuid"`
	Size int `json:"size"`
}

type Armor struct {
	DamageBiochemical int `json:"damage_biochemical"`
	DamageEnergy int `json:"damage_energy"`
	DamageStun int `json:"damage_stun"`
	SignalElectromagnetic float64 `json:"signal_electromagnetic"`
	DamagePhysical float64 `json:"damage_physical"`
	SignalCrossSection float64 `json:"signal_cross_section"`
	SignalInfrared float64 `json:"signal_infrared"`
	DamageDistortion int `json:"damage_distortion"`
	DamageThermal int `json:"damage_thermal"`
}

type Children struct {
	DisplayName string `json:"display_name"`
	DamageMax int `json:"damage_max"`
	Item *Item `json:"item"`
	ClassName string `json:"class_name"`
	Type string `json:"type"`
	MinSize any `json:"min_size"`
	Name string `json:"name"`
	Children []any `json:"children"`
	Health int `json:"health"`
	Position any `json:"position"`
	MaxSize any `json:"max_size"`
	SubType string `json:"sub_type"`
}

type CompatibleTypes struct {
	Type string `json:"type"`
	SubTypes []string `json:"sub_types"`
}

type Components struct {
	Name string `json:"name"`
	Category string `json:"category"`
	Manufacturer string `json:"manufacturer"`
	ComponentSize string `json:"component_size"`
	Size string `json:"size"`
	Mounts int `json:"mounts"`
	Details string `json:"details"`
	Quantity int `json:"quantity"`
	ComponentClass string `json:"component_class"`
	Type string `json:"type"`
}

type Cooler struct {
	CoolingRate int `json:"cooling_rate"`
	SuppressionIrFactor float64 `json:"suppression_ir_factor"`
	SuppressionHeatFactor float64 `json:"suppression_heat_factor"`
}

type CounterMeasure struct {
	Damages []any `json:"damages"`
	Type any `json:"type"`
	Capacity int `json:"capacity"`
	DamagePerShot int `json:"damage_per_shot"`
	Modes []*Modes `json:"modes"`
	Class any `json:"class"`
	Range int `json:"range"`
	Regeneration any `json:"regeneration"`
	Ammunition *Ammunition `json:"ammunition"`
}

type Crew struct {
	Min int `json:"min"`
	Max any `json:"max"`
	Weapon int `json:"weapon"`
	Operation any `json:"operation"`
}

type DamageFalloffs struct {
	MinDistance *MinDistance `json:"min_distance"`
	PerMeter *PerMeter `json:"per_meter"`
	MinDamage *MinDamage `json:"min_damage"`
}

type Damages struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Damage float64 `json:"damage"`
}

type Data struct {
	Fuel *Fuel `json:"fuel"`
	Description string `json:"description"`
	Slug string `json:"slug"`
	Quantum *Quantum `json:"quantum"`
	Id int `json:"id"`
	ShieldHp int `json:"shield_hp"`
	ClassName string `json:"class_name"`
	Foci []string `json:"foci"`
	Emission *Emission `json:"emission"`
	PersonalInventory int `json:"personal_inventory"`
	Skus []*Skus `json:"skus"`
	Parts []*Parts `json:"parts"`
	Components []*Components `json:"components"`
	PledgeUrl string `json:"pledge_url"`
	Hardpoints []*Hardpoints `json:"hardpoints"`
	VehicleInventory float64 `json:"vehicle_inventory"`
	UpdatedAt string `json:"updated_at"`
	Crew *Crew `json:"crew"`
	SizeClass int `json:"size_class"`
	Insurance *Insurance `json:"insurance"`
	Size string `json:"size"`
	Mass int `json:"mass"`
	ProductionStatus string `json:"production_status"`
	ChassisId int `json:"chassis_id"`
	Msrp int `json:"msrp"`
	CargoCapacity int `json:"cargo_capacity"`
	Version string `json:"version"`
	Manufacturer *Manufacturer `json:"manufacturer"`
	Shops []*Shops `json:"shops"`
	Name string `json:"name"`
	Speed *Speed `json:"speed"`
	Armor *Armor `json:"armor"`
	Loaner []any `json:"loaner"`
	Sizes *Sizes `json:"sizes"`
	Uuid string `json:"uuid"`
	Health int `json:"health"`
	Type string `json:"type"`
	ProductionNote string `json:"production_note"`
	Agility *Agility `json:"agility"`
}

type Emission struct {
	EmIdle int `json:"em_idle"`
	EmMax int `json:"em_max"`
	Ir int `json:"ir"`
}

type Emp struct {
	DamageBiochemical int `json:"damage_biochemical"`
	SignalElectromagnetic float64 `json:"signal_electromagnetic"`
	DamageDistortion int `json:"damage_distortion"`
	DamageThermal int `json:"damage_thermal"`
	SignalInfrared float64 `json:"signal_infrared"`
	SignalCrossSection float64 `json:"signal_cross_section"`
	DamagePhysical float64 `json:"damage_physical"`
	DamageEnergy int `json:"damage_energy"`
	DamageStun int `json:"damage_stun"`
}

type FlightController struct {
	ScmSpeed int `json:"scm_speed"`
	MaxSpeed int `json:"max_speed"`
	Pitch int `json:"pitch"`
	Yaw int `json:"yaw"`
	Roll int `json:"roll"`
}

type Fuel struct {
	IntakeRate int `json:"intake_rate"`
	Usage *Usage `json:"usage"`
	Capacity int `json:"capacity"`
}

type FuelIntake struct {
	MinimumRate float64 `json:"minimum_rate"`
	FuelPushRate int `json:"fuel_push_rate"`
}

type FuelTank struct {
	FillRate int `json:"fill_rate"`
	DrainRate int `json:"drain_rate"`
	Capacity float64 `json:"capacity"`
}

type Hardpoints struct {
	Type string `json:"type"`
	Children []*Children `json:"children"`
	MinSize int `json:"min_size"`
	Health any `json:"health"`
	SubType string `json:"sub_type"`
	Name string `json:"name"`
	Position any `json:"position"`
	ClassName string `json:"class_name"`
	Item *Item `json:"item"`
	MaxSize int `json:"max_size"`
}

type Insurance struct {
	ExpediteCost int `json:"expedite_cost"`
	ClaimTime float64 `json:"claim_time"`
	ExpediteTime float64 `json:"expedite_time"`
}

type Inventory struct {
	Length int `json:"length"`
	Dimension int `json:"dimension"`
	Scu float64 `json:"scu"`
	ScuConverted float64 `json:"scu_converted"`
	Unit string `json:"unit"`
	Width int `json:"width"`
	Height int `json:"height"`
}

type Item struct {
	SubType string `json:"sub_type"`
	Emp *Emp `json:"emp"`
	Thruster *Thruster `json:"thruster"`
	FuelTank *FuelTank `json:"fuel_tank"`
	MinSize int `json:"min_size"`
	Link string `json:"link"`
	Cooler *Cooler `json:"cooler"`
	Name string `json:"name"`
	Manufacturer *Manufacturer `json:"manufacturer"`
	Ports []*Ports `json:"ports"`
	FlightController *FlightController `json:"flight_controller"`
	MaxSize int `json:"max_size"`
	MaxMissiles int `json:"max_missiles"`
	SelfDestruct *SelfDestruct `json:"self_destruct"`
	Size int `json:"size"`
	ClassName string `json:"class_name"`
	Type string `json:"type"`
	Mass int `json:"mass"`
	PowerPlant *PowerPlant `json:"power_plant"`
	Grade string `json:"grade"`
	Class string `json:"class"`
	QuantumDrive *QuantumDrive `json:"quantum_drive"`
	UpdatedAt string `json:"updated_at"`
	CounterMeasure *CounterMeasure `json:"counter_measure"`
	Inventory *Inventory `json:"inventory"`
	Uuid string `json:"uuid"`
	MaxMounts int `json:"max_mounts"`
	FuelIntake *FuelIntake `json:"fuel_intake"`
	Shield *Shield `json:"shield"`
	Version string `json:"version"`
}

type Items struct {
	OptimalInventory int `json:"optimal_inventory"`
	Type string `json:"type"`
	Version string `json:"version"`
	SubType string `json:"sub_type"`
	PriceRange *PriceRange `json:"price_range"`
	MaxInventory int `json:"max_inventory"`
	Name string `json:"name"`
	Buyable bool `json:"buyable"`
	Sellable bool `json:"sellable"`
	PriceCalculated float64 `json:"price_calculated"`
	AutoConsume bool `json:"auto_consume"`
	Rentable bool `json:"rentable"`
	MaxPremium int `json:"max_premium"`
	Uuid string `json:"uuid"`
	Inventory int `json:"inventory"`
	RefreshRate int `json:"refresh_rate"`
	BasePrice float64 `json:"base_price"`
	AutoRestock bool `json:"auto_restock"`
	BasePriceOffset int `json:"base_price_offset"`
	MaxDiscount int `json:"max_discount"`
}

type Manufacturer struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Meta struct {
	ValidRelations []string `json:"valid_relations"`
	ProcessedAt string `json:"processed_at"`
}

type MinDamage struct {
	Thermal int `json:"thermal"`
	Biochemical int `json:"biochemical"`
	Stun int `json:"stun"`
	Physical int `json:"physical"`
	Energy int `json:"energy"`
	Distortion int `json:"distortion"`
}

type MinDistance struct {
	Thermal int `json:"thermal"`
	Biochemical int `json:"biochemical"`
	Stun int `json:"stun"`
	Physical int `json:"physical"`
	Energy int `json:"energy"`
	Distortion int `json:"distortion"`
}

type Missile struct {
	LockRangeMin int `json:"lock_range_min"`
	TrackingSignalMin int `json:"tracking_signal_min"`
	ClusterSize int `json:"cluster_size"`
	SignalType string `json:"signal_type"`
	LockRangeMax int `json:"lock_range_max"`
	LockAngle int `json:"lock_angle"`
	Speed float64 `json:"speed"`
	Damages []*Damages `json:"damages"`
	LockTime int `json:"lock_time"`
	FuelTankSize int `json:"fuel_tank_size"`
	ExplosionRadiusMax int `json:"explosion_radius_max"`
	ExplosionRadiusMin int `json:"explosion_radius_min"`
	DamageTotal float64 `json:"damage_total"`
}

type Modes struct {
	PelletsPerShot int `json:"pellets_per_shot"`
	InterdictionEffectTime int `json:"interdiction_effect_time"`
	CalibrationProcessAngleLimit int `json:"calibration_process_angle_limit"`
	MinCalibrationRequirement int `json:"min_calibration_requirement"`
	SpoolUpTime int `json:"spool_up_time"`
	DriveSpeed int `json:"drive_speed"`
	Rpm int `json:"rpm"`
	AmmoPerShot int `json:"ammo_per_shot"`
	EngageSpeed int `json:"engage_speed"`
	StageOneAccelRate int `json:"stage_one_accel_rate"`
	CalibrationRate int `json:"calibration_rate"`
	MaxCalibrationRequirement int `json:"max_calibration_requirement"`
	StageTwoAccelRate int `json:"stage_two_accel_rate"`
	CalibrationWarningAngleLimit int `json:"calibration_warning_angle_limit"`
	Type string `json:"type"`
	DamagePerSecond float64 `json:"damage_per_second"`
	Mode string `json:"mode"`
	CooldownTime int `json:"cooldown_time"`
}

type Parts struct {
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	DamageMax int `json:"damage_max"`
	Children []*Children `json:"children"`
}

type PerMeter struct {
	Physical int `json:"physical"`
	Energy int `json:"energy"`
	Distortion int `json:"distortion"`
	Thermal int `json:"thermal"`
	Biochemical int `json:"biochemical"`
	Stun int `json:"stun"`
}

type Piercability struct {
	DamageFalloffLevel1 int `json:"damage_falloff_level_1"`
	DamageFalloffLevel2 int `json:"damage_falloff_level_2"`
	DamageFalloffLevel3 int `json:"damage_falloff_level_3"`
	MaxPenetrationThickness float64 `json:"max_penetration_thickness"`
}

type Ports struct {
	Sizes *Sizes `json:"sizes"`
	CompatibleTypes []*CompatibleTypes `json:"compatible_types"`
	Position string `json:"position"`
	Tags []any `json:"tags"`
	RequiredTags []any `json:"required_tags"`
	EquippedItem any `json:"equipped_item"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
}

type PowerPlant struct {
	PowerOutput float64 `json:"power_output"`
}

type PriceRange struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type Quantum struct {
	QuantumRange float64 `json:"quantum_range"`
	QuantumSpeed float64 `json:"quantum_speed"`
	QuantumSpoolTime int `json:"quantum_spool_time"`
	QuantumFuelCapacity float64 `json:"quantum_fuel_capacity"`
}

type QuantumDrive struct {
	Modes []*Modes `json:"modes"`
	QuantumFuelRequirement float64 `json:"quantum_fuel_requirement"`
	JumpRange string `json:"jump_range"`
	DisconnectRange int `json:"disconnect_range"`
	ThermalEnergyDraw *ThermalEnergyDraw `json:"thermal_energy_draw"`
}

type RegenDelay struct {
	Downed int `json:"downed"`
	Damage int `json:"damage"`
}

type Regeneration struct {
	RequestedAmmoLoad int `json:"requested_ammo_load"`
	Cooldown float64 `json:"cooldown"`
	CostPerBullet float64 `json:"cost_per_bullet"`
	RequestedRegenPerSec int `json:"requested_regen_per_sec"`
}

type SelfDestruct struct {
	Damage int `json:"damage"`
	Radius int `json:"radius"`
	MinRadius int `json:"min_radius"`
	PhysRadius int `json:"phys_radius"`
	MinPhysRadius int `json:"min_phys_radius"`
	Time int `json:"time"`
}

type Shield struct {
	MaxShieldHealth int `json:"max_shield_health"`
	MaxShieldRegen int `json:"max_shield_regen"`
	DecayRatio float64 `json:"decay_ratio"`
	RegenDelay *RegenDelay `json:"regen_delay"`
	MaxReallocation int `json:"max_reallocation"`
	ReallocationRate int `json:"reallocation_rate"`
}

type Shops struct {
	Name string `json:"name"`
	Position string `json:"position"`
	ProfitMargin int `json:"profit_margin"`
	Link string `json:"link"`
	Version string `json:"version"`
	Items []*Items `json:"items"`
	Uuid string `json:"uuid"`
	NameRaw string `json:"name_raw"`
}

type Sizes struct {
	Min int `json:"min"`
	Max int `json:"max"`
	Length float64 `json:"length"`
	Beam float64 `json:"beam"`
	Height int `json:"height"`
}

type Skus struct {
	Price int `json:"price"`
	Available int `json:"available"`
	ImportedAt string `json:"imported_at"`
	Title string `json:"title"`
}

type Speed struct {
	ScmToZero float64 `json:"scm_to_zero"`
	MaxToZero float64 `json:"max_to_zero"`
	Scm int `json:"scm"`
	Max int `json:"max"`
	ZeroToScm float64 `json:"zero_to_scm"`
	ZeroToMax float64 `json:"zero_to_max"`
}

type ThermalEnergyDraw struct {
	InFlight int `json:"in_flight"`
	RampDown int `json:"ramp_down"`
	PostRampDown int `json:"post_ramp_down"`
	PreRampUp int `json:"pre_ramp_up"`
	RampUp int `json:"ramp_up"`
}

type Thruster struct {
	ThrustCapacity float64 `json:"thrust_capacity"`
	MinHealthThrustMultiplier float64 `json:"min_health_thrust_multiplier"`
	FuelBurnPer10kNewton float64 `json:"fuel_burn_per_10k_newton"`
	Type string `json:"type"`
}

type Usage struct {
	Vtol int `json:"vtol"`
	Main int `json:"main"`
	Maneuvering int `json:"maneuvering"`
	Retro int `json:"retro"`
}

type VehicleWeapon struct {
	DamagePerShot float64 `json:"damage_per_shot"`
	Damages []*Damages `json:"damages"`
	Regeneration *Regeneration `json:"regeneration"`
	Class any `json:"class"`
	Type string `json:"type"`
	Capacity int `json:"capacity"`
	Ammunition *Ammunition `json:"ammunition"`
	Modes []*Modes `json:"modes"`
	Range int `json:"range"`
}


