package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

type datatable struct {
	resultData          interface{}
	cardinalityFiltered int64
	cardinality         int64
	requestValues       Query
	columns             []string
	collection          *mongo.Collection
	data                []map[string]interface{}
	filterSearch        map[string]interface{}
}

func NewDatatable(collection *mongo.Collection, requestValues Query, columns []string) *datatable {
	return &datatable{
		collection:    collection,
		requestValues: requestValues,
		columns:       columns,
	}
}

func (d *datatable) generate() {
	count, err := d.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return
	}

	d.cardinality = count

	// SEARCH MASIH BUG //LEBIH BAIK PAKAI INDEXING
	sSearch := d.requestValues.SSearch
	if sSearch != "" {
		var agg []interface{}
		for _, column := range d.columns {
			agg = append(agg, map[string]interface{}{column: map[string]string{"$regex": sSearch, "$options": "i"}})
		}
		d.filterSearch = map[string]interface{}{"$or": agg}

		//indexing
		//d.filterSearch = map[string]interface{}{
		//	"$text": map[string]interface{}{"$search":""},
		//}
	}

	countColFilt, err := d.collection.CountDocuments(context.TODO(), d.filterSearch)
	if err != nil {
		log.Println(err)
		return
	}

	d.cardinalityFiltered = countColFilt
	//sort data
	sortedData := d.customSort()
	skip := d.requestValues.IDisplayStart
	limit := d.requestValues.IDisplayLength

	opt := options.Find()
	opt.Skip = &skip
	opt.Limit = &limit
	opt.SetProjection(bson.M{"_id": 0})
	opt.SetSort(bson.M{sortedData["column_name"].(string): sortedData["sort"]})

	cur, err := d.collection.Find(context.TODO(), d.filterSearch, opt)
	if err != nil {
		log.Println(err)
		return
	}

	var r []map[string]interface{}
	for cur.Next(context.TODO()) {
		var temp map[string]interface{}
		err = cur.Decode(&temp)
		if err != nil {
			continue
		}
		r = append(r, temp)
	}

	d.data = r
}

func (d *datatable) customSort() map[string]interface{} {
	if d.requestValues.ISortCol0 != "" && d.requestValues.ISortingCols > 0 {
		index, err := strconv.Atoi(d.requestValues.ISortCol0)
		if err != nil {
			return map[string]interface{}{}
		}
		columnName := d.columns[index]
		return map[string]interface{}{"column_name": columnName, "sort": isReverse(d.requestValues.SSortDir0)}
	}
	return nil
}

func isReverse(strDirection string) int {
	if strDirection == "desc" {
		return -1
	} else {
		return 1
	}
}

func (d *datatable) result() Data {
	data := Data{
		SEcho:                d.requestValues.SEcho,
		ITotalRecords:        d.cardinality,
		ITotalDisplayRecords: d.cardinalityFiltered,
		Data:                 []string{},
	}

	if d.data != nil {
		data.Data = d.data
	}
	return data
}
