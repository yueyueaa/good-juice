// Code generated by ent, DO NOT EDIT.

package videometadatum

import (
	"juice/app/public/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldID, id))
}

// VideoID applies equality check predicate on the "video_id" field. It's identical to VideoIDEQ.
func VideoID(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldUserID, v))
}

// CoverURL applies equality check predicate on the "cover_url" field. It's identical to CoverURLEQ.
func CoverURL(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldCoverURL, v))
}

// VideoURL applies equality check predicate on the "video_url" field. It's identical to VideoURLEQ.
func VideoURL(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoURL, v))
}

// VideoIntro applies equality check predicate on the "video_intro" field. It's identical to VideoIntroEQ.
func VideoIntro(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoIntro, v))
}

// VideoType applies equality check predicate on the "video_type" field. It's identical to VideoTypeEQ.
func VideoType(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoType, v))
}

// PublishAddress applies equality check predicate on the "publish_address" field. It's identical to PublishAddressEQ.
func PublishAddress(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldPublishAddress, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldUpdateTime, v))
}

// VideoIDEQ applies the EQ predicate on the "video_id" field.
func VideoIDEQ(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoID, v))
}

// VideoIDNEQ applies the NEQ predicate on the "video_id" field.
func VideoIDNEQ(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldVideoID, v))
}

// VideoIDIn applies the In predicate on the "video_id" field.
func VideoIDIn(vs ...int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldVideoID, vs...))
}

// VideoIDNotIn applies the NotIn predicate on the "video_id" field.
func VideoIDNotIn(vs ...int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldVideoID, vs...))
}

// VideoIDGT applies the GT predicate on the "video_id" field.
func VideoIDGT(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldVideoID, v))
}

// VideoIDGTE applies the GTE predicate on the "video_id" field.
func VideoIDGTE(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldVideoID, v))
}

// VideoIDLT applies the LT predicate on the "video_id" field.
func VideoIDLT(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldVideoID, v))
}

// VideoIDLTE applies the LTE predicate on the "video_id" field.
func VideoIDLTE(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldVideoID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldUserID, v))
}

// CoverURLEQ applies the EQ predicate on the "cover_url" field.
func CoverURLEQ(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldCoverURL, v))
}

// CoverURLNEQ applies the NEQ predicate on the "cover_url" field.
func CoverURLNEQ(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldCoverURL, v))
}

// CoverURLIn applies the In predicate on the "cover_url" field.
func CoverURLIn(vs ...string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldCoverURL, vs...))
}

// CoverURLNotIn applies the NotIn predicate on the "cover_url" field.
func CoverURLNotIn(vs ...string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldCoverURL, vs...))
}

// CoverURLGT applies the GT predicate on the "cover_url" field.
func CoverURLGT(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldCoverURL, v))
}

// CoverURLGTE applies the GTE predicate on the "cover_url" field.
func CoverURLGTE(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldCoverURL, v))
}

// CoverURLLT applies the LT predicate on the "cover_url" field.
func CoverURLLT(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldCoverURL, v))
}

// CoverURLLTE applies the LTE predicate on the "cover_url" field.
func CoverURLLTE(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldCoverURL, v))
}

// CoverURLContains applies the Contains predicate on the "cover_url" field.
func CoverURLContains(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldContains(FieldCoverURL, v))
}

// CoverURLHasPrefix applies the HasPrefix predicate on the "cover_url" field.
func CoverURLHasPrefix(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldHasPrefix(FieldCoverURL, v))
}

// CoverURLHasSuffix applies the HasSuffix predicate on the "cover_url" field.
func CoverURLHasSuffix(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldHasSuffix(FieldCoverURL, v))
}

// CoverURLEqualFold applies the EqualFold predicate on the "cover_url" field.
func CoverURLEqualFold(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEqualFold(FieldCoverURL, v))
}

// CoverURLContainsFold applies the ContainsFold predicate on the "cover_url" field.
func CoverURLContainsFold(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldContainsFold(FieldCoverURL, v))
}

// VideoURLEQ applies the EQ predicate on the "video_url" field.
func VideoURLEQ(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoURL, v))
}

// VideoURLNEQ applies the NEQ predicate on the "video_url" field.
func VideoURLNEQ(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldVideoURL, v))
}

// VideoURLIn applies the In predicate on the "video_url" field.
func VideoURLIn(vs ...string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldVideoURL, vs...))
}

// VideoURLNotIn applies the NotIn predicate on the "video_url" field.
func VideoURLNotIn(vs ...string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldVideoURL, vs...))
}

// VideoURLGT applies the GT predicate on the "video_url" field.
func VideoURLGT(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldVideoURL, v))
}

// VideoURLGTE applies the GTE predicate on the "video_url" field.
func VideoURLGTE(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldVideoURL, v))
}

// VideoURLLT applies the LT predicate on the "video_url" field.
func VideoURLLT(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldVideoURL, v))
}

// VideoURLLTE applies the LTE predicate on the "video_url" field.
func VideoURLLTE(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldVideoURL, v))
}

// VideoURLContains applies the Contains predicate on the "video_url" field.
func VideoURLContains(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldContains(FieldVideoURL, v))
}

// VideoURLHasPrefix applies the HasPrefix predicate on the "video_url" field.
func VideoURLHasPrefix(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldHasPrefix(FieldVideoURL, v))
}

// VideoURLHasSuffix applies the HasSuffix predicate on the "video_url" field.
func VideoURLHasSuffix(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldHasSuffix(FieldVideoURL, v))
}

// VideoURLEqualFold applies the EqualFold predicate on the "video_url" field.
func VideoURLEqualFold(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEqualFold(FieldVideoURL, v))
}

// VideoURLContainsFold applies the ContainsFold predicate on the "video_url" field.
func VideoURLContainsFold(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldContainsFold(FieldVideoURL, v))
}

// VideoIntroEQ applies the EQ predicate on the "video_intro" field.
func VideoIntroEQ(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoIntro, v))
}

// VideoIntroNEQ applies the NEQ predicate on the "video_intro" field.
func VideoIntroNEQ(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldVideoIntro, v))
}

// VideoIntroIn applies the In predicate on the "video_intro" field.
func VideoIntroIn(vs ...string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldVideoIntro, vs...))
}

// VideoIntroNotIn applies the NotIn predicate on the "video_intro" field.
func VideoIntroNotIn(vs ...string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldVideoIntro, vs...))
}

// VideoIntroGT applies the GT predicate on the "video_intro" field.
func VideoIntroGT(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldVideoIntro, v))
}

// VideoIntroGTE applies the GTE predicate on the "video_intro" field.
func VideoIntroGTE(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldVideoIntro, v))
}

// VideoIntroLT applies the LT predicate on the "video_intro" field.
func VideoIntroLT(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldVideoIntro, v))
}

// VideoIntroLTE applies the LTE predicate on the "video_intro" field.
func VideoIntroLTE(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldVideoIntro, v))
}

// VideoIntroContains applies the Contains predicate on the "video_intro" field.
func VideoIntroContains(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldContains(FieldVideoIntro, v))
}

// VideoIntroHasPrefix applies the HasPrefix predicate on the "video_intro" field.
func VideoIntroHasPrefix(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldHasPrefix(FieldVideoIntro, v))
}

// VideoIntroHasSuffix applies the HasSuffix predicate on the "video_intro" field.
func VideoIntroHasSuffix(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldHasSuffix(FieldVideoIntro, v))
}

// VideoIntroIsNil applies the IsNil predicate on the "video_intro" field.
func VideoIntroIsNil() predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIsNull(FieldVideoIntro))
}

// VideoIntroNotNil applies the NotNil predicate on the "video_intro" field.
func VideoIntroNotNil() predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotNull(FieldVideoIntro))
}

// VideoIntroEqualFold applies the EqualFold predicate on the "video_intro" field.
func VideoIntroEqualFold(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEqualFold(FieldVideoIntro, v))
}

// VideoIntroContainsFold applies the ContainsFold predicate on the "video_intro" field.
func VideoIntroContainsFold(v string) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldContainsFold(FieldVideoIntro, v))
}

// VideoTypeEQ applies the EQ predicate on the "video_type" field.
func VideoTypeEQ(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldVideoType, v))
}

// VideoTypeNEQ applies the NEQ predicate on the "video_type" field.
func VideoTypeNEQ(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldVideoType, v))
}

// VideoTypeIn applies the In predicate on the "video_type" field.
func VideoTypeIn(vs ...int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldVideoType, vs...))
}

// VideoTypeNotIn applies the NotIn predicate on the "video_type" field.
func VideoTypeNotIn(vs ...int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldVideoType, vs...))
}

// VideoTypeGT applies the GT predicate on the "video_type" field.
func VideoTypeGT(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldVideoType, v))
}

// VideoTypeGTE applies the GTE predicate on the "video_type" field.
func VideoTypeGTE(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldVideoType, v))
}

// VideoTypeLT applies the LT predicate on the "video_type" field.
func VideoTypeLT(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldVideoType, v))
}

// VideoTypeLTE applies the LTE predicate on the "video_type" field.
func VideoTypeLTE(v int) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldVideoType, v))
}

// PublishAddressEQ applies the EQ predicate on the "publish_address" field.
func PublishAddressEQ(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldPublishAddress, v))
}

// PublishAddressNEQ applies the NEQ predicate on the "publish_address" field.
func PublishAddressNEQ(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldPublishAddress, v))
}

// PublishAddressIn applies the In predicate on the "publish_address" field.
func PublishAddressIn(vs ...int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldPublishAddress, vs...))
}

// PublishAddressNotIn applies the NotIn predicate on the "publish_address" field.
func PublishAddressNotIn(vs ...int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldPublishAddress, vs...))
}

// PublishAddressGT applies the GT predicate on the "publish_address" field.
func PublishAddressGT(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldPublishAddress, v))
}

// PublishAddressGTE applies the GTE predicate on the "publish_address" field.
func PublishAddressGTE(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldPublishAddress, v))
}

// PublishAddressLT applies the LT predicate on the "publish_address" field.
func PublishAddressLT(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldPublishAddress, v))
}

// PublishAddressLTE applies the LTE predicate on the "publish_address" field.
func PublishAddressLTE(v int32) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldPublishAddress, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.FieldNotNull(FieldUpdateTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VideoMetadatum) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VideoMetadatum) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VideoMetadatum) predicate.VideoMetadatum {
	return predicate.VideoMetadatum(sql.NotPredicates(p))
}
