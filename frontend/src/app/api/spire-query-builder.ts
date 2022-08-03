import util from "util";

type QueryBuilderRequest = {
  select: string;
  where: string;
  whereOr: string;
  orderBy: string;
  orderDirection: string;
  groupBy: string;
  includes: string;
  limit: number;
  page: number;
}

export class SpireQueryBuilder {
  private selects: string[]       = [];
  private wheres: string[]        = [];
  private whereOrs: string[]      = [];
  private orderBys: string[]      = [];
  private groupBys: string[]      = [];
  private includesParam: string[] = [];
  private orderDirections: string = "";
  private limitParam: number      = 1000;
  private pageParam: number       = 0;

  translateOperator(operator) {
    switch (operator) {
      case "&":
        return "_bitwiseand_"
      case "notlike":
        return "_notlike_"
      case "like":
        return "_like_"
      case "!=":
        return "_ne_"
      case "=":
        return "__"
      case "<":
        return "_lt_"
      case "<=":
        return "_lte_"
      case ">":
        return "_gt_"
      case ">=":
        return "_gte_"
      default:
        return ""
    }
  }

  getFilterCount() {
    let filterCount: number = 0

    if (Object.keys(this.wheres).length > 0) {
      filterCount += Object.keys(this.wheres).length
    }

    if (Object.keys(this.whereOrs).length > 0) {
      filterCount += Object.keys(this.whereOrs).length
    }

    return filterCount
  }

  select(fields: string[]) {
    this.selects = fields

    return this
  }

  groupBy(fields: string[]) {
    this.groupBys = fields

    return this
  }

  orderBy(fields: string[]) {
    this.orderBys = fields

    return this
  }

  orderDirection(direction: string) {
    this.orderDirections = direction

    return this
  }

  includes(includes: string[]) {
    this.includesParam = includes

    return this
  }

  where(field, operator, value) {
    const where = util.format(
      "%s%s%s",
      field,
      this.translateOperator(operator),
      value
    )

    this.wheres.push(where)

    return this
  }

  whereOr(field, operator, value) {
    const where = util.format(
      "%s%s%s",
      field,
      this.translateOperator(operator),
      value
    )

    this.whereOrs.push(where)

    return this
  }

  limit(limit: number) {
    this.limitParam = limit

    return this
  }

  page(page: number) {
    this.pageParam = page

    return this
  }

  get() {
    let request = {} as QueryBuilderRequest;
    if (Object.keys(this.wheres).length > 0) {
      request.where = this.wheres.join(".")
    }
    if (Object.keys(this.whereOrs).length > 0) {
      request.whereOr = this.whereOrs.join(".")
    }
    if (Object.keys(this.selects).length > 0) {
      request.select = this.selects.join(".")
    }
    if (Object.keys(this.orderBys).length > 0) {
      request.orderBy = this.orderBys.join(".")
    }
    if (Object.keys(this.groupBys).length > 0) {
      request.groupBy = this.groupBys.join(".")
    }
    if (Object.keys(this.includesParam).length > 0) {
      request.includes = this.includesParam.join(",")
    }
    if (this.orderDirections) {
      request.orderDirection = this.orderDirections
    }
    console.log(this.pageParam)

    if (this.pageParam > 0) {
      request.page = (this.pageParam - 1)
    }

    console.log(request)

    request.limit = this.limitParam

    return request
  }
}
