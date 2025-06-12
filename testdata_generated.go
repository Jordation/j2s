package j2s

type Skus struct {
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	Available  float64 `json:"available"`
	ImportedAt string  `json:"imported_at"`
}

type Cooler struct {
	SuppressionHeatFactor float64 `json:"suppression_heat_factor"`
	CoolingRate           float64 `json:"cooling_rate"`
	SuppressionIrFactor   float64 `json:"suppression_ir_factor"`
}

type VehicleWeapon struct {
	Regeneration  Regeneration `json:"regeneration"`
	Ammunition    Ammunition   `json:"ammunition"`
	Range         float64      `json:"range"`
	Class         any          `json:"class"`
	Type          string       `json:"type"`
	DamagePerShot float64      `json:"damage_per_shot"`
	Modes         []*Modes     `json:"modes"`
	Capacity      float64      `json:"capacity"`
	Damages       []*Damages   `json:"damages"`
}

type Ammunition struct {
	Uuid           string         `json:"uuid"`
	Size           float64        `json:"size"`
	Lifetime       float64        `json:"lifetime"`
	Speed          float64        `json:"speed"`
	Range          float64        `json:"range"`
	Piercability   Piercability   `json:"piercability"`
	DamageFalloffs DamageFalloffs `json:"damage_falloffs"`
}

type Damages struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Damage float64 `json:"damage"`
}

type Hardpoints struct {
	ClassName string      `json:"class_name"`
	Health    any         `json:"health"`
	Item      Item        `json:"item"`
	Name      string      `json:"name"`
	Position  any         `json:"position"`
	MinSize   float64     `json:"min_size"`
	Children  []*Children `json:"children"`
	MaxSize   float64     `json:"max_size"`
	Type      string      `json:"type"`
	SubType   string      `json:"sub_type"`
}

type Parts struct {
	Children    []*Children `json:"children"`
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	DamageMax   float64     `json:"damage_max"`
}

type Emp struct {
	SignalElectromagnetic float64 `json:"signal_electromagnetic"`
	SignalCrossSection    float64 `json:"signal_cross_section"`
	DamagePhysical        float64 `json:"damage_physical"`
	DamageDistortion      float64 `json:"damage_distortion"`
	DamageStun            float64 `json:"damage_stun"`
	SignalInfrared        float64 `json:"signal_infrared"`
	DamageThermal         float64 `json:"damage_thermal"`
	DamageBiochemical     float64 `json:"damage_biochemical"`
	DamageEnergy          float64 `json:"damage_energy"`
}

type Modes struct {
	DriveSpeed                   float64 `json:"drive_speed"`
	StageOneAccelRate            float64 `json:"stage_one_accel_rate"`
	EngageSpeed                  float64 `json:"engage_speed"`
	CalibrationRate              float64 `json:"calibration_rate"`
	CalibrationWarningAngleLimit float64 `json:"calibration_warning_angle_limit"`
	DamagePerSecond              float64 `json:"damage_per_second"`
	PelletsPerShot               float64 `json:"pellets_per_shot"`
	Rpm                          float64 `json:"rpm"`
	CalibrationProcessAngleLimit float64 `json:"calibration_process_angle_limit"`
	MinCalibrationRequirement    float64 `json:"min_calibration_requirement"`
	Type                         string  `json:"type"`
	CooldownTime                 float64 `json:"cooldown_time"`
	AmmoPerShot                  float64 `json:"ammo_per_shot"`
	MaxCalibrationRequirement    float64 `json:"max_calibration_requirement"`
	SpoolUpTime                  float64 `json:"spool_up_time"`
	StageTwoAccelRate            float64 `json:"stage_two_accel_rate"`
	InterdictionEffectTime       float64 `json:"interdiction_effect_time"`
	Mode                         string  `json:"mode"`
}

type CounterMeasure struct {
	Class         any        `json:"class"`
	Type          any        `json:"type"`
	Ammunition    Ammunition `json:"ammunition"`
	Range         float64    `json:"range"`
	Modes         []*Modes   `json:"modes"`
	Capacity      float64    `json:"capacity"`
	DamagePerShot float64    `json:"damage_per_shot"`
	Damages       []any      `json:"damages"`
	Regeneration  any        `json:"regeneration"`
}

type Quantum struct {
	QuantumSpeed        float64 `json:"quantum_speed"`
	QuantumSpoolTime    float64 `json:"quantum_spool_time"`
	QuantumFuelCapacity float64 `json:"quantum_fuel_capacity"`
	QuantumRange        float64 `json:"quantum_range"`
}

type Speed struct {
	ZeroToScm float64 `json:"zero_to_scm"`
	ZeroToMax float64 `json:"zero_to_max"`
	ScmToZero float64 `json:"scm_to_zero"`
	MaxToZero float64 `json:"max_to_zero"`
	Scm       float64 `json:"scm"`
	Max       float64 `json:"max"`
}

type DamageFalloffs struct {
	PerMeter    PerMeter    `json:"per_meter"`
	MinDamage   MinDamage   `json:"min_damage"`
	MinDistance MinDistance `json:"min_distance"`
}

type Meta struct {
	ProcessedAt    string   `json:"processed_at"`
	ValidRelations []string `json:"valid_relations"`
}

type Shops struct {
	Items        []*Items `json:"items"`
	Uuid         string   `json:"uuid"`
	NameRaw      string   `json:"name_raw"`
	Name         string   `json:"name"`
	Position     string   `json:"position"`
	ProfitMargin float64  `json:"profit_margin"`
	Link         string   `json:"link"`
	Version      string   `json:"version"`
}

type PerMeter struct {
	Physical    float64 `json:"physical"`
	Energy      float64 `json:"energy"`
	Distortion  float64 `json:"distortion"`
	Thermal     float64 `json:"thermal"`
	Biochemical float64 `json:"biochemical"`
	Stun        float64 `json:"stun"`
}

type Fuel struct {
	Capacity   float64 `json:"capacity"`
	IntakeRate float64 `json:"intake_rate"`
	Usage      Usage   `json:"usage"`
}

type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Components struct {
	Mounts         float64 `json:"mounts"`
	ComponentSize  string  `json:"component_size"`
	Category       string  `json:"category"`
	Type           string  `json:"type"`
	Name           string  `json:"name"`
	Manufacturer   string  `json:"manufacturer"`
	ComponentClass string  `json:"component_class"`
	Details        string  `json:"details"`
	Size           string  `json:"size"`
	Quantity       float64 `json:"quantity"`
}

type Manufacturer struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Link string `json:"link"`
}

type Data struct {
	Quantum           Quantum       `json:"quantum"`
	UpdatedAt         string        `json:"updated_at"`
	Fuel              Fuel          `json:"fuel"`
	Manufacturer      Manufacturer  `json:"manufacturer"`
	ProductionNote    string        `json:"production_note"`
	Name              string        `json:"name"`
	ProductionStatus  string        `json:"production_status"`
	Description       string        `json:"description"`
	Msrp              float64       `json:"msrp"`
	Slug              string        `json:"slug"`
	Components        []*Components `json:"components"`
	Skus              []*Skus       `json:"skus"`
	CargoCapacity     float64       `json:"cargo_capacity"`
	Id                float64       `json:"id"`
	ClassName         string        `json:"class_name"`
	Emission          Emission      `json:"emission"`
	Mass              float64       `json:"mass"`
	Uuid              string        `json:"uuid"`
	Speed             Speed         `json:"speed"`
	Hardpoints        []*Hardpoints `json:"hardpoints"`
	Parts             []*Parts      `json:"parts"`
	PledgeUrl         string        `json:"pledge_url"`
	Sizes             Sizes         `json:"sizes"`
	Version           string        `json:"version"`
	PersonalInventory float64       `json:"personal_inventory"`
	Crew              Crew          `json:"crew"`
	Size              string        `json:"size"`
	Loaner            []any         `json:"loaner"`
	ShieldHp          float64       `json:"shield_hp"`
	Insurance         Insurance     `json:"insurance"`
	Armor             Armor         `json:"armor"`
	VehicleInventory  float64       `json:"vehicle_inventory"`
	SizeClass         float64       `json:"size_class"`
	Agility           Agility       `json:"agility"`
	Type              string        `json:"type"`
	Health            float64       `json:"health"`
	Shops             []*Shops      `json:"shops"`
	Foci              []string      `json:"foci"`
	ChassisId         float64       `json:"chassis_id"`
}

type Usage struct {
	Retro       float64 `json:"retro"`
	Vtol        float64 `json:"vtol"`
	Main        float64 `json:"main"`
	Maneuvering float64 `json:"maneuvering"`
}

type Items struct {
	AutoConsume      bool       `json:"auto_consume"`
	Uuid             string     `json:"uuid"`
	MaxInventory     float64    `json:"max_inventory"`
	Rentable         bool       `json:"rentable"`
	OptimalInventory float64    `json:"optimal_inventory"`
	AutoRestock      bool       `json:"auto_restock"`
	BasePrice        float64    `json:"base_price"`
	MaxDiscount      float64    `json:"max_discount"`
	Type             string     `json:"type"`
	MaxPremium       float64    `json:"max_premium"`
	Inventory        float64    `json:"inventory"`
	RefreshRate      float64    `json:"refresh_rate"`
	Sellable         bool       `json:"sellable"`
	PriceRange       PriceRange `json:"price_range"`
	Name             string     `json:"name"`
	SubType          string     `json:"sub_type"`
	BasePriceOffset  float64    `json:"base_price_offset"`
	Buyable          bool       `json:"buyable"`
	Version          string     `json:"version"`
	PriceCalculated  float64    `json:"price_calculated"`
}

type MinDamage struct {
	Stun        float64 `json:"stun"`
	Physical    float64 `json:"physical"`
	Energy      float64 `json:"energy"`
	Distortion  float64 `json:"distortion"`
	Thermal     float64 `json:"thermal"`
	Biochemical float64 `json:"biochemical"`
}

type Regeneration struct {
	RequestedRegenPerSec float64 `json:"requested_regen_per_sec"`
	RequestedAmmoLoad    float64 `json:"requested_ammo_load"`
	Cooldown             float64 `json:"cooldown"`
	CostPerBullet        float64 `json:"cost_per_bullet"`
}

type Ship struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

type Missile struct {
	ExplosionRadiusMax float64    `json:"explosion_radius_max"`
	DamageTotal        float64    `json:"damage_total"`
	ClusterSize        float64    `json:"cluster_size"`
	SignalType         string     `json:"signal_type"`
	Speed              float64    `json:"speed"`
	ExplosionRadiusMin float64    `json:"explosion_radius_min"`
	LockTime           float64    `json:"lock_time"`
	LockRangeMax       float64    `json:"lock_range_max"`
	TrackingSignalMin  float64    `json:"tracking_signal_min"`
	FuelTankSize       float64    `json:"fuel_tank_size"`
	LockRangeMin       float64    `json:"lock_range_min"`
	LockAngle          float64    `json:"lock_angle"`
	Damages            []*Damages `json:"damages"`
}

type Piercability struct {
	DamageFalloffLevel1     float64 `json:"damage_falloff_level_1"`
	DamageFalloffLevel2     float64 `json:"damage_falloff_level_2"`
	DamageFalloffLevel3     float64 `json:"damage_falloff_level_3"`
	MaxPenetrationThickness float64 `json:"max_penetration_thickness"`
}

type Insurance struct {
	ExpediteCost float64 `json:"expedite_cost"`
	ClaimTime    float64 `json:"claim_time"`
	ExpediteTime float64 `json:"expedite_time"`
}

type SelfDestruct struct {
	Damage        float64 `json:"damage"`
	Radius        float64 `json:"radius"`
	MinRadius     float64 `json:"min_radius"`
	PhysRadius    float64 `json:"phys_radius"`
	MinPhysRadius float64 `json:"min_phys_radius"`
	Time          float64 `json:"time"`
}

type PowerPlant struct {
	PowerOutput float64 `json:"power_output"`
}

type MinDistance struct {
	Physical    float64 `json:"physical"`
	Energy      float64 `json:"energy"`
	Distortion  float64 `json:"distortion"`
	Thermal     float64 `json:"thermal"`
	Biochemical float64 `json:"biochemical"`
	Stun        float64 `json:"stun"`
}

type Sizes struct {
	Beam   float64 `json:"beam"`
	Height float64 `json:"height"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	Length float64 `json:"length"`
}

type FuelIntake struct {
	FuelPushRate float64 `json:"fuel_push_rate"`
	MinimumRate  float64 `json:"minimum_rate"`
}

type Agility struct {
	Acceleration Acceleration `json:"acceleration"`
	Pitch        float64      `json:"pitch"`
	Yaw          float64      `json:"yaw"`
	Roll         float64      `json:"roll"`
}

type Inventory struct {
	Height       float64 `json:"height"`
	Length       float64 `json:"length"`
	Dimension    float64 `json:"dimension"`
	Scu          float64 `json:"scu"`
	ScuConverted float64 `json:"scu_converted"`
	Unit         string  `json:"unit"`
	Width        float64 `json:"width"`
}

type Acceleration struct {
	Vtol         float64 `json:"vtol"`
	Maneuvering  float64 `json:"maneuvering"`
	MainG        float64 `json:"main_g"`
	RetroG       float64 `json:"retro_g"`
	VtolG        float64 `json:"vtol_g"`
	ManeuveringG float64 `json:"maneuvering_g"`
	Main         float64 `json:"main"`
	Retro        float64 `json:"retro"`
}

type Children struct {
	DisplayName string  `json:"display_name"`
	MaxSize     any     `json:"max_size"`
	SubType     string  `json:"sub_type"`
	Health      float64 `json:"health"`
	Type        string  `json:"type"`
	Item        Item    `json:"item"`
	Name        string  `json:"name"`
	DamageMax   float64 `json:"damage_max"`
	Children    []any   `json:"children"`
	Position    any     `json:"position"`
	ClassName   string  `json:"class_name"`
	MinSize     any     `json:"min_size"`
}

type Ports struct {
	CompatibleTypes []*CompatibleTypes `json:"compatible_types"`
	Tags            []any              `json:"tags"`
	RequiredTags    []any              `json:"required_tags"`
	EquippedItem    any                `json:"equipped_item"`
	Name            string             `json:"name"`
	DisplayName     string             `json:"display_name"`
	Position        string             `json:"position"`
	Sizes           Sizes              `json:"sizes"`
}

type RegenDelay struct {
	Damage float64 `json:"damage"`
	Downed float64 `json:"downed"`
}

type FlightController struct {
	Roll     float64 `json:"roll"`
	ScmSpeed float64 `json:"scm_speed"`
	MaxSpeed float64 `json:"max_speed"`
	Pitch    float64 `json:"pitch"`
	Yaw      float64 `json:"yaw"`
}

type FuelTank struct {
	Capacity  float64 `json:"capacity"`
	FillRate  float64 `json:"fill_rate"`
	DrainRate float64 `json:"drain_rate"`
}

type ThermalEnergyDraw struct {
	PreRampUp    float64 `json:"pre_ramp_up"`
	RampUp       float64 `json:"ramp_up"`
	InFlight     float64 `json:"in_flight"`
	RampDown     float64 `json:"ramp_down"`
	PostRampDown float64 `json:"post_ramp_down"`
}

type Shield struct {
	DecayRatio       float64    `json:"decay_ratio"`
	RegenDelay       RegenDelay `json:"regen_delay"`
	MaxReallocation  float64    `json:"max_reallocation"`
	ReallocationRate float64    `json:"reallocation_rate"`
	MaxShieldHealth  float64    `json:"max_shield_health"`
	MaxShieldRegen   float64    `json:"max_shield_regen"`
}

type QuantumDrive struct {
	QuantumFuelRequirement float64           `json:"quantum_fuel_requirement"`
	JumpRange              string            `json:"jump_range"`
	DisconnectRange        float64           `json:"disconnect_range"`
	ThermalEnergyDraw      ThermalEnergyDraw `json:"thermal_energy_draw"`
	Modes                  []*Modes          `json:"modes"`
}

type Emission struct {
	Ir     float64 `json:"ir"`
	EmIdle float64 `json:"em_idle"`
	EmMax  float64 `json:"em_max"`
}

type Armor struct {
	DamagePhysical        float64 `json:"damage_physical"`
	DamageBiochemical     float64 `json:"damage_biochemical"`
	DamageDistortion      float64 `json:"damage_distortion"`
	SignalInfrared        float64 `json:"signal_infrared"`
	SignalElectromagnetic float64 `json:"signal_electromagnetic"`
	DamageThermal         float64 `json:"damage_thermal"`
	DamageStun            float64 `json:"damage_stun"`
	SignalCrossSection    float64 `json:"signal_cross_section"`
	DamageEnergy          float64 `json:"damage_energy"`
}

type Crew struct {
	Max       any     `json:"max"`
	Weapon    float64 `json:"weapon"`
	Operation any     `json:"operation"`
	Min       float64 `json:"min"`
}

type Item struct {
	Inventory      Inventory      `json:"inventory"`
	FuelIntake     FuelIntake     `json:"fuel_intake"`
	Class          string         `json:"class"`
	Version        string         `json:"version"`
	ClassName      string         `json:"class_name"`
	PowerPlant     PowerPlant     `json:"power_plant"`
	MaxMissiles    float64        `json:"max_missiles"`
	Link           string         `json:"link"`
	QuantumDrive   QuantumDrive   `json:"quantum_drive"`
	SubType        string         `json:"sub_type"`
	Uuid           string         `json:"uuid"`
	Size           float64        `json:"size"`
	Mass           float64        `json:"mass"`
	FuelTank       FuelTank       `json:"fuel_tank"`
	Name           string         `json:"name"`
	Ports          []*Ports       `json:"ports"`
	Type           string         `json:"type"`
	Cooler         Cooler         `json:"cooler"`
	MaxSize        float64        `json:"max_size"`
	Shield         Shield         `json:"shield"`
	MinSize        float64        `json:"min_size"`
	Manufacturer   Manufacturer   `json:"manufacturer"`
	Emp            Emp            `json:"emp"`
	CounterMeasure CounterMeasure `json:"counter_measure"`
	MaxMounts      float64        `json:"max_mounts"`
	UpdatedAt      string         `json:"updated_at"`
	Thruster       Thruster       `json:"thruster"`
	SelfDestruct   SelfDestruct   `json:"self_destruct"`
	VehicleWeapon  VehicleWeapon  `json:"vehicle_weapon"`
	Grade          string         `json:"grade"`
}

type Thruster struct {
	MinHealthThrustMultiplier float64 `json:"min_health_thrust_multiplier"`
	FuelBurnPer10kNewton      float64 `json:"fuel_burn_per_10k_newton"`
	Type                      string  `json:"type"`
	ThrustCapacity            float64 `json:"thrust_capacity"`
}

type CompatibleTypes struct {
	Type     string   `json:"type"`
	SubTypes []string `json:"sub_types"`
}
